// Repository is the Clean Architecture package for accessing DB.
package repository

import (
	"database/sql"
	"go-api-clean/entity"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	//ctx := context.Background()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/postsclean")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO poststable (Title, Text)VALUES(?,?)",
		post.Title, post.Text)
	if err != nil {
		log.Println(err.Error())
	}
	defer insert.Close()
	var insertedPost entity.Post
	for insert.Next() {
		err = insert.Scan(&insertedPost.ID, &insertedPost.Title, &insertedPost.Text)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &insertedPost, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	//ctx := context.Background()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/postsclean")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM poststable")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	var allPosts []entity.Post
	for results.Next() {
		var tempPost entity.Post
		err = results.Scan(&tempPost.ID, &tempPost.Title, &tempPost.Text)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		allPosts = append(allPosts, tempPost)
	}
	return allPosts, nil
}
