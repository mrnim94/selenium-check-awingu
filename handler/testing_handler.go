package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	yaml "gopkg.in/yaml.v2"
	"net/http"
	"selenium-check-awingu/helper/automate"
	"selenium-check-awingu/helper/git/git_impl"
	"selenium-check-awingu/log"
	"selenium-check-awingu/model"
	req "selenium-check-awingu/model/req"
	"selenium-check-awingu/model/testing"
	"selenium-check-awingu/repository"
	"strings"
)

type TestingHandler struct {
	TestingRepo repository.TestingRepo
	Automate    automate.Automate
}

// RunTesting godoc
// @Summary Chạy testing bằng tên Job
// @Tags testing-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestTesting true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /tester/run-testing [post]
func (t *TestingHandler) HandlerRunTesting(c echo.Context) error {
	req := req.RequestTesting{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("bắt đầu sử lý request JobName: " + req.JobName)
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho JobName: " + req.JobName + " thành công")
	}

	jobTesting, err := t.TestingRepo.SelectJobByName(c.Request().Context(), req.JobName)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("lấy thông tin của JobName: " + req.JobName + " trong database thành công")
	}

	jobUsers, err := t.TestingRepo.SelectAllUserByJobId(c.Request().Context(), jobTesting.JobId)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	infoGitHub, err := t.TestingRepo.SelectGithubByJobId(c.Request().Context(), jobTesting.JobId)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("lấy thông tin Github của JobName: " + req.JobName + " trong database thành công")
	}

	authGithub := &git_impl.ConfigurationGithub{
		GithubAccessToken: infoGitHub.AccessToken,
	}
	result := git_impl.NewGitHubConnection(authGithub)

	contentScriptTesting, err := result.GetContentYaml(infoGitHub.Owner, infoGitHub.Repo, infoGitHub.Path)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	yml := &testing.YamlTesting{}
	err = yaml.Unmarshal(contentScriptTesting, &yml)
	if err != nil {
		log.Error(err.Error() + "Check lại file Yaml của bạn")
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//check Hooks
	if len(yml.Hooks) == 0 {
		patchSplit := strings.Split(infoGitHub.Path, "/")
		contentBeginTest := string(contentScriptTesting) + "\nhooks:"
		for _, action := range yml.Actions {
			pathGithub := patchSplit[0] + "/actions/" + action + "/tasks/main.yml"
			contentActionScript, err := result.GetContentYaml(infoGitHub.Owner, infoGitHub.Repo, pathGithub)
			if err != nil {
				log.Error(err.Error())
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}

			lines := strings.Split(string(contentActionScript), "\n")
			for _, line := range lines {
				if line != "hooks:" {
					contentBeginTest = contentBeginTest + "\n" + line
				}
			}
		}
		contentMergeActionScript := []byte(contentBeginTest)
		err = yaml.Unmarshal(contentMergeActionScript, &yml)
		if err != nil {
			log.Error(err.Error() + " Check lại file Yaml folder actions của bạn")
			return c.JSON(http.StatusInternalServerError, model.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
		}
	}

	for _, user := range jobUsers {
		go t.Automate.RobotAutoImpl(*yml, user, t.TestingRepo, jobTesting.JobId)
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

// AddJob godoc
// @Summary Thêm Job cho tool testing
// @Tags testing-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddJob true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /tester/add-job [post]
func (t *TestingHandler) HandlerAddJob(c echo.Context) error {
	req := req.RequestAddJob{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("bắt đầu sử lý request AddJob: " + req.JobName)
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho AddJob: " + req.JobName + " thành công")
	}

	jobId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	jobTest := model.JobsTesting{
		JobId: jobId.String(),
		JobName:   req.JobName,
		Status:    req.Status,
	}

	jobTest, err = t.TestingRepo.SaveJobTest(c.Request().Context(), jobTest)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       jobTest,
	})
}

// DeleteJob godoc
// @Summary Xóa Job cho tool testing
// @Tags testing-service
// @Accept  json
// @Produce  json
// @Param jobID path string true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /tester/delete-job/{jobID} [delete]
func (t *TestingHandler) HandlerDeleteJob(c echo.Context) error {
	jobID := c.Param("jobid")
	log.Info("Thực hiện Delete Job: " + jobID)
	err := t.TestingRepo.RemoveJobTest(c.Request().Context(), jobID)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       nil,
	})
}

// AddGithubJob godoc
// @Summary Thêm thông tin Github cho tool testing
// @Tags testing-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddGithubJob true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /tester/add-github [post]
func (t *TestingHandler) HandlerAddGithubJob(c echo.Context) error {
	req := req.RequestAddGithubJob{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("bắt đầu sử lý request AddGithub: " + req.Repo)
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho AddGithub: " + req.Repo + " thành công")
	}

	jobGithub := model.JobsGithub{
		AccessToken: req.AccessToken,
		Owner:       req.Owner,
		Repo:        req.Repo,
		Path:        req.Path,
		JobID:       req.JobID,
	}

	jobGithub, err := t.TestingRepo.SaveJobGithub(c.Request().Context(), jobGithub)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       jobGithub,
	})
}

// AddUserJob godoc
// @Summary Thêm thông tin User sử dụng cho bài test
// @Tags testing-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddUserJob true "testing"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 409 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /tester/add-user [post]
func (t *TestingHandler) HandlerAddUserJob(c echo.Context) error {
	req := req.RequestAddUserJob{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("bắt đầu sử lý request AddUser: " + req.Username)
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		log.Info("kiểm trả các thông số gửi lên cho AddUser: " + req.Username + " thành công")
	}

	jobUser := model.JobsUser{
		Username:  req.Username,
		Password:  req.Password,
		JobId:     req.JobId,
		Status:    req.Status,
	}

	jobUser, err := t.TestingRepo.SaveJobUser(c.Request().Context(), jobUser)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       jobUser,
	})
}