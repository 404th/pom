package loader

import (
	"fmt"
	"time"
)

func Load(tm int) {
	// Define the loader animation frames
	frames := []string{
		"[           ]",
		"[🍅         ]",
		"[🍅🍅       ]",
		"[🍅🍅🍅     ]",
		"[ 🍅🍅🍅    ]",
		"[  🍅🍅🍅   ]",
		"[   🍅🍅🍅  ]",
		"[    🍅🍅🍅 ]",
		"[     🍅🍅🍅]",
		"[       🍅🍅]",
		"[         🍅]",
	}

	fmt.Println("⏰ Time get started...")

	// Loop through the frames to create the loader animation
	for i := 0; i < tm*60*10; i++ {
		// Print the current frame
		fmt.Print("\r" + frames[i%len(frames)])
		time.Sleep(100 * time.Millisecond) // Delay between frames
	}
}
