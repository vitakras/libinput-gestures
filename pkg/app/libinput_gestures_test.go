package app_test

import (
	"errors"

	"github.com/vitakras/libinput-gestures/pkg/libinput"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"

	. "github.com/vitakras/libinput-gestures/pkg/app"
	appfakes "github.com/vitakras/libinput-gestures/pkg/app/appfakes"
)

var _ = Describe("LibinputGestures", func() {
	var (
		fakeProcessor     *appfakes.FakeProcessor
		fakeDebugStreamer *appfakes.FakeDebugStreamer
		libinputGestures  *LibinputGestures
	)

	BeforeEach(func() {
		fakeProcessor = new(appfakes.FakeProcessor)
		fakeDebugStreamer = new(appfakes.FakeDebugStreamer)
		libinputGestures = NewLibinputGestures(fakeProcessor, fakeDebugStreamer)
	})

	Context("ProcessEvents", func() {
		Context("DebugStreamer.Start() returns an error", func() {
			It("returns an error", func() {
				fakeDebugStreamer.StartReturns(errors.New("Stream Failed to Start"))

				err := libinputGestures.ProcessEvents()
				Expect(err).To(HaveOccurred())
				Expect(fakeDebugStreamer.StartCallCount()).To(Equal(1))
			})
		})

		Context("DebugStreamer.Close() returns true", func() {
			It("has not errors", func() {
				fakeDebugStreamer.ClosedReturns(true)

				err := libinputGestures.ProcessEvents()
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeDebugStreamer.ClosedCallCount()).To(Equal(1))
			})
		})

		Context("Processor.Process() return an error", func() {
			It("returns an error", func() {
				fakeDebugStreamer.ClosedReturns(false)
				fakeProcessor.ProcessReturns(errors.New("Failed to Process DebugEvent"))

				err := libinputGestures.ProcessEvents()
				Expect(err).To(HaveOccurred())
				Expect(fakeDebugStreamer.ClosedCallCount()).To(Equal(1))
				Expect(fakeProcessor.ProcessCallCount()).To(Equal(1))
			})
		})

		Context("It Processes DebugEvents successfully", func() {
			It("has not errors", func() {
				debugEvent := new(libinput.DebugEvent)

				fakeDebugStreamer.ClosedReturnsOnCall(0, false)
				fakeDebugStreamer.ClosedReturnsOnCall(1, true)
				fakeDebugStreamer.ReadReturns(debugEvent)

				err := libinputGestures.ProcessEvents()
				Expect(err).NotTo(HaveOccurred())
				Expect(fakeDebugStreamer.ClosedCallCount()).To(Equal(2))
				Expect(fakeDebugStreamer.ReadCallCount()).To(Equal(1))

				expectedDebugEvent := fakeProcessor.ProcessArgsForCall(0)
				Expect(debugEvent).To(Equal(expectedDebugEvent))
			})
		})
	})
})
