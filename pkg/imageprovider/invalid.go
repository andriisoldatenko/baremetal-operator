package imageprovider

import (
	"fmt"
)

type ImageBuildInvalidError struct {
	err error
}

func (ibf ImageBuildInvalidError) Error() string {
	return fmt.Sprintf("Cannot generate image: %s", ibf.err.Error())
}

func (ibf ImageBuildInvalidError) Unwrap() error {
	return ibf.err
}

func BuildInvalidError(err error) ImageBuildInvalidError {
	return ImageBuildInvalidError{err: err}
}

type ImageNotReadyError struct{}

func (inr ImageNotReadyError) Error() string {
	return "Image is not ready yet"
}
