package requests

import (
	"github.com/tylerb/codecs"
	"github.com/tylerb/codecs/services"
)

var codecService services.CodecService

// Codecs returns the
// "github.com/tylerb/codecs/services".CodecService currently in use
// by this library.
func Codecs() services.CodecService {
	if codecService == nil {
		codecService = services.NewWebCodecService()
	}
	return codecService
}

// SetCodecs can be used to change the
// "github.com/tylerb/codecs/services".CodecService used by this
// library.
func SetCodecs(newService services.CodecService) {
	codecService = newService
}

// AddCodec adds a "github.com/tylerb/codecs".Codec to the
// "github.com/tylerb/codecs/services".CodecService currently in use
// by this library.
func AddCodec(codec codecs.Codec) {
	Codecs().AddCodec(codec)
}
