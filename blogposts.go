package blogposts

import (
	"io/fs"
)

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
