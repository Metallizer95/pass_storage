package passportctrl

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"store_server/internal/usecase/errs"
	"store_server/internal/usecase/passport"
	"strconv"
)

func (ctrl *controller) savePassport(r passport.Model, c *gin.Context) {
	pass := ctrl.SaveUseCase.Save(r)
	if pass == nil {
		errResponse := errs.NewErrModel(errs.ErrObjectAlreadyExists)
		c.XML(http.StatusOK, errResponse)
		return
	}

	c.XML(http.StatusOK, pass)
	ctrl.logger.Info("return statusOK")
}

func (ctrl *controller) uploadFile(c *gin.Context) ([]*zip.File, error) {
	file, _, err := c.Request.FormFile("data")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	msgLength, err := strconv.Atoi(c.GetHeader("Content-Length"))
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(file, int64(msgLength))
	if err != nil {
		return nil, err
	}

	return zipReader.File, nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
