package hold

import (
	"testing"
	"time"

	"github.com/nu7hatch/gouuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHoldChan(t *testing.T) {
	Convey("set then get", t, func() {
		id, _ := uuid.NewV4()
		Hold.Set(id, "derp")
		derp := Hold.Get(id)
		So(derp, ShouldEqual, "derp")
	})
	Convey("get then set", t, func() {
		id, _ := uuid.NewV4()
		go func() {
			time.Sleep(time.Second / 10)
			Hold.Set(id, "derpsy")
		}()
		derp := Hold.Get(id)
		So(derp, ShouldEqual, "derpsy")
	})
}