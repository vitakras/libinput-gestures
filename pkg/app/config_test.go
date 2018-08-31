package app_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/vitakras/libinput-gestures/pkg/app"
	. "github.com/vitakras/libinput-gestures/pkg/app/appfakes"
)

var _ = Describe("Config", func() {
	Context("LoadConfigFromFile", func() {
		It("Can load config from a file", func() {
			config, err := LoadConfigFromFile("test/config.yml")

			Expect(err).NotTo(HaveOccurred())
			Expect(config.Swipe[3].Left.Command).To(Equal("xdotool key alt+Right"))
			Expect(config.Swipe[3].Right.Command).To(Equal("xdotool key alt+Left"))
			Expect(config.Swipe[3].Up.Command).To(Equal("xdotool key super"))
			Expect(config.Swipe[3].Down.Command).To(Equal("xdotool key super"))

			Expect(config.Pinch.Out.Command).To(Equal("xdotool key ctrl+minus"))
			Expect(config.Pinch.In.Command).To(Equal("xdotool key ctrl+plus"))

			Expect(config.Threshold.Pinch).To(Equal(float32(0.4)))
			Expect(config.Threshold.Swipe).To(Equal(float32(0.4)))

			Expect(config.Interval.Pinch).To(Equal(float32(0.1)))
			Expect(config.Interval.Swipe).To(Equal(float32(0.8)))
		})
	})

	Context("GetCommand", func() {
		var (
			config  *Config
			gesture *FakeGesture
		)

		BeforeEach(func() {
			config, _ = LoadConfigFromFile("test/config.yml")
		})

		Context("When the gesture is not complete", func() {
			It("returns nil", func() {
				gesture = new(FakeGesture)
				gesture.IsCompleteReturns(false)

				command := config.GetCommand(gesture)
				Expect(command).To(BeNil())
				Expect(gesture.IsCompleteCallCount()).To(Equal(1))
			})
		})

		Context("When a command is not found", func() {
			It("returns nil", func() {
				gesture = new(FakeGesture)
				gesture.IsCompleteReturns(true)
				gesture.FingersReturns(5)
				gesture.TypeReturns("unknown")

				command := config.GetCommand(gesture)
				Expect(command).To(BeNil())
			})
		})

		Context("When the gesture is swipe", func() {
			Context("and the gesture is found", func() {
				It("returns the command", func() {
					gesture = new(FakeGesture)
					gesture.IsCompleteReturns(true)
					gesture.FingersReturns(3)
					gesture.TypeReturns("swipe")
					gesture.DirectionReturns("left")

					command := config.GetCommand(gesture)
					Expect(command.Command).To(Equal("xdotool key alt+Right"))
				})
			})
		})

		Context("When the gesture is pinch", func() {
			Context("and the gesture is found", func() {
				It("returns the command", func() {
					gesture = new(FakeGesture)
					gesture.IsCompleteReturns(true)
					gesture.FingersReturns(3)
					gesture.TypeReturns("pinch")
					gesture.DirectionReturns("out")

					command := config.GetCommand(gesture)
					Expect(command.Command).To(Equal("xdotool key ctrl+minus"))
				})
			})
		})
	})
})
