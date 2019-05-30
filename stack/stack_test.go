package stack

import (
	"testing"
	. "github.com/smartystreets/goconvey1/convey"
)

func TestConstructor(t *testing.T) {
	Convey("queue method test", t, func() {

		Convey("construct test", func() {
			So(Constructor(), ShouldNotBeNil)
		})

		q := Constructor()
		q.Push(10)

		Convey("push test", func() {
			So(q.Pop(), ShouldEqual, 10)
		})

		/*q.Push(12)
		q.Push(20)
		Convey("peek test", func() {
			So(q.Peek(), ShouldEqual, 20)
		})*/

		Convey("empty test", func() {
			So(q.Empty(), ShouldEqual, false)
		})
	})
}
