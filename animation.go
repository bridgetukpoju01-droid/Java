Question 5: ASCII Art Animation System


Build a system that creates simple ASCII art animations by generating frames programmatically. The system should support creating animations with transitions, rotations, and simple movements.

Requirements:

Create Animation struct with:

NewAnimation(text string, frames int) constructor

GenerateSpinFrames() creates rotation effect

GenerateWaveFrames() creates wave effect

GenerateZoomFrames() creates zoom in/out effect

GetFrame(index int) string returns specific frame

Play() string returns all frames with delay markers

Each frame should be exactly 10 lines tall

Animations must loop seamlessly

All ASCII art generated programmatically (no external files)

File Structure:
ascii-art/
├── main.go
├── animation/
│   ├── animation.go
│   └── animation_test.go
└── effects/
    ├── spin.go
    ├── wave.go
    ├── zoom.go
    └── effects_test.go


