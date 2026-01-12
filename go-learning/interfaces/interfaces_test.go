package interfaces

import (
	"math"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "radius 1 gives area Pi",
			radius: 1,
			want:   math.Pi,
		},
		{
			name:   "radius 5 gives area 25*Pi",
			radius: 5,
			want:   25 * math.Pi,
		},
		{
			name:   "radius 0 gives area 0",
			radius: 0,
			want:   0,
		},
		{
			name:   "radius 10 gives area 100*Pi",
			radius: 10,
			want:   100 * math.Pi,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{Radius: tt.radius}
			if got := c.Area(); got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "radius 1 gives perimeter 2*Pi",
			radius: 1,
			want:   2 * math.Pi,
		},
		{
			name:   "radius 5 gives perimeter 10*Pi",
			radius: 5,
			want:   10 * math.Pi,
		},
		{
			name:   "radius 0 gives perimeter 0",
			radius: 0,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{Radius: tt.radius}
			if got := c.Perimeter(); got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{
			name:   "4x6 rectangle has area 24",
			width:  4,
			height: 6,
			want:   24,
		},
		{
			name:   "square 5x5 has area 25",
			width:  5,
			height: 5,
			want:   25,
		},
		{
			name:   "zero width gives area 0",
			width:  0,
			height: 10,
			want:   0,
		},
		{
			name:   "1x1 has area 1",
			width:  1,
			height: 1,
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{Width: tt.width, Height: tt.height}
			if got := r.Area(); got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name   string
		width  float64
		height float64
		want   float64
	}{
		{
			name:   "4x6 rectangle has perimeter 20",
			width:  4,
			height: 6,
			want:   20,
		},
		{
			name:   "square 5x5 has perimeter 20",
			width:  5,
			height: 5,
			want:   20,
		},
		{
			name:   "1x1 has perimeter 4",
			width:  1,
			height: 1,
			want:   4,
		},
		{
			name:   "10x2 has perimeter 24",
			width:  10,
			height: 2,
			want:   24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{Width: tt.width, Height: tt.height}
			if got := r.Perimeter(); got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintShapeInfo(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
	}{
		{
			name:  "prints circle info",
			shape: Circle{Radius: 5},
		},
		{
			name:  "prints rectangle info",
			shape: Rectangle{Width: 4, Height: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// PrintShapeInfo just prints, no return value to check
			// This test ensures it doesn't panic
			PrintShapeInfo(tt.shape)
		})
	}
}

func TestConsoleWriter_Write(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantErr bool
	}{
		{
			name:    "writes simple string",
			data:    "Hello, World!",
			wantErr: false,
		},
		{
			name:    "writes empty string",
			data:    "",
			wantErr: false,
		},
		{
			name:    "writes long string",
			data:    "This is a longer test message for the console writer",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cw := ConsoleWriter{}
			if err := cw.Write(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("ConsoleWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileWriter_Write(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		data     string
		wantErr  bool
	}{
		{
			name:     "writes to output.txt",
			filename: "output.txt",
			data:     "Hello, File!",
			wantErr:  false,
		},
		{
			name:     "writes to log file",
			filename: "app.log",
			data:     "Log entry",
			wantErr:  false,
		},
		{
			name:     "writes empty data",
			filename: "empty.txt",
			data:     "",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fw := FileWriter{Filename: tt.filename}
			if err := fw.Write(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("FileWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestShape_Interface verifies that Circle and Rectangle implement Shape
func TestShape_Interface(t *testing.T) {
	var _ Shape = Circle{}
	var _ Shape = Rectangle{}
}

// TestWriter_Interface verifies that ConsoleWriter and FileWriter implement Writer
func TestWriter_Interface(t *testing.T) {
	var _ Writer = ConsoleWriter{}
	var _ Writer = FileWriter{}
}

func TestRunInterfaces(t *testing.T) {
	// Smoke test - ensures function runs without panic
	t.Run("runs without panic", func(t *testing.T) {
		RunInterfaces()
	})
}

func Test_demonstratePolymorphism(t *testing.T) {
	t.Run("runs without panic", func(t *testing.T) {
		demonstratePolymorphism()
	})
}

func Test_demonstrateInterfaceTypes(t *testing.T) {
	t.Run("runs without panic", func(t *testing.T) {
		demonstrateInterfaceTypes()
	})
}

func Test_demonstrateEmptyInterface(t *testing.T) {
	t.Run("runs without panic", func(t *testing.T) {
		demonstrateEmptyInterface()
	})
}
