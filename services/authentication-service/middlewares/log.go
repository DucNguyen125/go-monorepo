package middlewares

import (
	"context"
	"fmt"
	"math"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func Logger() grpc.UnaryServerInterceptor {
	return func(context context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		service := path.Dir(info.FullMethod)[1:]
		method := path.Base(info.FullMethod)
		rawIp, _ := peer.FromContext(context)
		clientIP := rawIp.Addr.String()
		start := time.Now()
		response, err := handler(context, request)
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		logDetail := log.Fields{
			"service":  service,
			"latency":  fmt.Sprintf("%dns", latency), // time to process
			"clientIP": clientIP,
			"method":   method,
			"code":     status.Code(err).String(),
			"data":     request,
		}
		logger := log.WithFields(logDetail)

		if err != nil {
			logger.Error(err)
		} else {
			logger.Info("[GRPC]")
		}
		return response, err
	}
}
