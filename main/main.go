package main

import (
	"alnoor/blogposts/reader"
	"alnoor/blogposts/renderer"
	"fmt"
	"net/http"
	"os"
)

func main() {
	PostGenerator()
	fileServer()
}

func fileServer() {
	fs := http.FileServer(http.Dir("dest"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}

func PostGenerator() {
	// read examples dir for posts as markdown
	posts, err := reader.NewPostsFromFS(os.DirFS("../examples"))
	if err != nil {
		fmt.Print(err)
		panic("could not read dir 'examples'")
	}

	// init new renderer
	postRenderer, err := renderer.NewPostRenderer()
	if err != nil {
		panic("could not read make new post renderer")
	}

	makeDest()

	// prepare file mode to write, create if not exist or clear if exist
	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	// index file
	file, _ := os.OpenFile("dest/index.html", mode, 0644)
	// close in end of func
	defer file.Close()
	// write template to index file
	postRenderer.RenderIndex(file, posts)
	fmt.Printf("finish write to file %s\n", file.Name())

	// iterate throw posts
	for _, post := range posts {
		// open file or create one for each post
		file, _ := os.OpenFile(fmt.Sprintf("dest/%s.html", post.SanitisedTitle()), mode, 0644)
		defer file.Close()
		// render post file
		postRenderer.Render(file, post)
		fmt.Printf("finish write to file %s\n", file.Name())
	}
}

func makeDest() {
	// Directory path you want to create
	dirPath := "dest"

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory if it does not exist
		err := os.Mkdir(dirPath, 0755) // 0755 is the Unix permission mode for the directory
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Directory created:", dirPath)
	} else if err != nil {
		// Handle other errors
		fmt.Println("Error:", err)
		return
	}
}
