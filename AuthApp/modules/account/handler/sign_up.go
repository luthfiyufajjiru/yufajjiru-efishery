package handler

import (
	"authapp/adapters"
	"authapp/modules/account/dal"
	am "authapp/modules/account/models"
	"authapp/modules/account/services"
	st "authapp/modules/token/services"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SignUp(x *adapters.Adapters) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var u am.UserDTO

		err := ctx.BodyParser(&u)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		} else if err == nil && (!services.PhoneNumberValidation(u.PhoneNumber) || u.Name == "") {
			return fiber.NewError(http.StatusBadRequest, "invalid input")
		}

		if u.Password == "" {
			u.Password = st.DefaultPassword()
		}

		err = x.SQL.DB.Transaction(func(tx *gorm.DB) error {
			if errs := dal.CreateUser(tx, u.MapToUserModel()); errs != nil {
				return errs
			}
			return nil
		})

		if err != nil {
			s := err.Error()

			pattern := `\.(.*)$`
			mtch, _ := regexp.Compile(pattern)

			xx := mtch.FindAllString(s, -1)

			x := err.Error()

			if k := strings.Contains(s, "UNIQUE"); k {
				s = "already exist"
			}

			if len(xx) > 0 {
				x = xx[0]
				x = strings.Replace(x, ".", "", -1)
				x = strings.Replace(x, "_", " ", -1)
				x = fmt.Sprintf("%s is %s", x, s)
			}

			return fiber.NewError(http.StatusBadRequest, x)
		}

		ctx.Status(http.StatusOK).SendString(u.Password)

		return err
	}
}
