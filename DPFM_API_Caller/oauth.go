package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-instagram-initial-auth-requests-rmq/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-instagram-initial-auth-requests-rmq/DPFM_API_Output_Formatter"
	"data-platform-api-instagram-initial-auth-requests-rmq/config"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) InstagramInitialAuth(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.InstagramInitialAuth {
	var instagramInitialAuth []dpfm_api_output_formatter.InstagramInitialAuth

	instagramInitialAuth = append(instagramInitialAuth, dpfm_api_output_formatter.InstagramInitialAuth{
		URL: conf.OAuth.AuthURL,
	})

	return &instagramInitialAuth
}
