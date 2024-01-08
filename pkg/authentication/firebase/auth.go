package auth

import (
	"context"
	"errors"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/byeol-i/firebase-auth-module/pkg/cache"
	"github.com/byeol-i/firebase-auth-module/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)


type FirebaseApp struct {
	app *firebase.App
	cache *cache.Cache
}

func NewFirebaseApp(path, projectId string) (*FirebaseApp, error) {
	// path := configManager.GetFirebaseCredFilePath()

	opt := option.WithCredentialsFile(path)
	firebaseConfig := &firebase.Config{ProjectID: projectId}

	app, err := firebase.NewApp(context.Background(), firebaseConfig, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	idTokenCache := cache.NewCache(5*time.Minute)

	return &FirebaseApp{
		app: app,
		cache: idTokenCache,
	}, nil
}

type GetResult struct {
	Result interface{}
	Error  error
}

func GetUserFromFirebase(app *FirebaseApp, ctx context.Context, uid string) GetResult {
	client, err := app.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return GetResult{
			Result: nil,
			Error:  errors.New("error getting Auth client"),
		}
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		logger.Error("error getting user", zap.Error(err), zap.String("uid", uid))
		return GetResult{
			Result: nil,
			Error:  errors.New("error getting user"),
		}
	}

	return GetResult{
		Result: u,
		Error:  nil,
	}
}

func VerifyIDTokenFromFirebase(app *FirebaseApp, ctx context.Context, idToken string, cache *cache.Cache) GetResult {
	client, err := app.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return GetResult{
			Result: nil,
			Error:  errors.New("error getting Auth client"),
		}
	}
	
	res, found := cache.Get(idToken)
	if found {
		return GetResult{
			Result: res,
			Error:  nil,
		}
	} else {
		decodedToken, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			logger.Error("Can't verify token", zap.Error(err))
			return GetResult{
				Result: nil,
				Error:  err,
			}
		}

		cache.Set(idToken, decodedToken.UID)

		return GetResult{
			Result: decodedToken.UID,
			Error:  nil,
		}
	}
}

func (hdl *FirebaseApp) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	result := make(chan GetResult)

	go func() {
		result <- GetUserFromFirebase(hdl, ctx, uid)
	}()
	select {
	case <-time.After(5 * time.Second):
		return nil, errors.New("timed out")
	case result := <-result:
		if result.Error != nil {
			return nil, result.Error
		}

		if _, ok := result.Result.(*auth.UserRecord); ok {
			return nil, errors.New("Type error!")
		}

		return result.Result.(*auth.UserRecord), nil
	}
}

func (hdl *FirebaseApp) CreateCustomToken(ctx context.Context, uid string) (string, error) {
	client, err := hdl.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return "", err
	}
	token, err := client.CustomToken(ctx, uid)

	if err != nil {
		logger.Error("Can't create custom token!", zap.Error(err))
		return "", err
	}

	return token, nil
}

func (hdl *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	result := make(chan GetResult)

	go func() {
		result <- VerifyIDTokenFromFirebase(hdl, ctx, idToken, hdl.cache)
	}()
	select {
	case <-time.After(5 * time.Second):
		return "", errors.New("timed out")
	case result := <-result:
		if result.Error != nil {
			return "", result.Error
		}

		if _, ok := result.Result.(string); ok {
			if (!ok) {
				return "", errors.New("Type error")
			}
		}

		return result.Result.(string), nil
	}
}
