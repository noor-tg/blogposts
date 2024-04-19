package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	// read dir files
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post

	// iterate throw dir files slice
	for _, f := range dir {
		// get each file content as Post struct
		post, err := getPost(filesystem, f.Name())

		if err != nil {
			return nil, err
		}
		// return err when needed
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, fileName string) (Post, error) {
	// open file and get content
	postFile, err := filesystem.Open(fileName)

	if err != nil {
		return Post{}, err
	}

	// close file
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	// read file content until eof or err
	postData, err := io.ReadAll(postFile)

	if err != nil {
		return Post{}, err
	}

	// convert file content to string and get the chars after 7 char to end as title
	post := Post{Title: string(postData)[7:]}

	return post, nil
}
