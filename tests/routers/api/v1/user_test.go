package v1_test

import (
	"bbs_backend/entity"
	v1 "bbs_backend/routers/api/v1"
	. "bbs_backend/tests/mock/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	var app *gin.Engine
	var rec *httptest.ResponseRecorder
	var (
		mockCtrl *gomock.Controller
		e        *MockUserModel
	)
	var baseUrl string = "/api/v1/user/"
	gin.SetMode(gin.ReleaseMode)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		e = NewMockUserModel(mockCtrl)

		app = gin.New()
		userRoute := v1.NewUserRouter(e)
		userRoute.Setup(app.Group("/api"))

		rec = httptest.NewRecorder()
	})

	It("should return all users", func() {
		req, _ := http.NewRequest("GET", baseUrl, nil)
		currentTime := time.Now()
		user := entity.User{
			ID:        1,
			Nickname:  "Nickname",
			Email:     "user@test.com",
			Password:  "p455w0rd",
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		}
		e.EXPECT().GetAll().DoAndReturn(func() (*[]entity.User, error) {
			users := &[]entity.User{user}
			return users, nil
		})
		var actual, expected string
		jsonData, _ := json.MarshalIndent([]entity.User{user}, "", "    ")
		expected = string(jsonData)

		app.ServeHTTP(rec, req)
		actual = rec.Body.String()

		Expect(rec.Code).To(Equal(http.StatusOK))
		Expect(actual).To(Equal(expected))
	})
})
