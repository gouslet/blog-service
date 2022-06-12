/*
 * File: /pkg/limiter/limiter.go                                               *
 * Project: blog-service                                                       *
 * Created At: Wednesday, 2022/06/8 , 09:41:44                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/10 , 06:18:02                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(c *gin.Context) string // get name of the 
	GetBucket(key string) (*ratelimit.Bucket, bool) // get a token bucket with name of key
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // add token buckets into the limiter
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
