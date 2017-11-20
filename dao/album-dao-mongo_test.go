package dao_test

import (
	"github.com/satori/go.uuid"
	"go-backend-sample/dao"
	"go-backend-sample/model"
	"testing"
)

func TestAlbumDAOMongo(t *testing.T) {
	authorDao, albumDao, err := dao.GetDAO(dao.MongoDAO, "")
	if err != nil {
		t.Error(err)
	}

	authorToSave := model.Author{
		Id:        uuid.NewV4().String(),
		Firstname: "Test1",
		Lastname:  "Test2",
	}

	authorSaved, err := authorDao.Upsert(&authorToSave)
	if err != nil {
		t.Error(err)
	}

	song1ToSave := model.Song{
		Title:  "Test1",
		Number: "1",
	}
	song2ToSave := model.Song{
		Title:  "Test2",
		Number: "2",
	}

	var songsToSave []model.Song
	songsToSave = append(songsToSave, song1ToSave)
	songsToSave = append(songsToSave, song2ToSave)

	albumToSave := model.Album{
		Id:          uuid.NewV4().String(),
		Title:       "Test",
		Description: "Description Test",
		AuthorId:    authorSaved.Id,
		Songs:       songsToSave,
	}

	albumSaved, err := albumDao.Upsert(&albumToSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("album saved", albumSaved)

	albums, err := albumDao.GetAll()
	if err != nil {
		t.Error(err)
	}

	t.Log("album found all", albums[0])

	oneAlbum, err := albumDao.Get(albums[0].Id)
	if err != nil {
		t.Error(err)
	}

	t.Log("album found one", oneAlbum)

	oneAlbum.Title = "Test2"
	oneAlbum.Description = "Description Test2"
	chg, err := albumDao.Upsert(oneAlbum)
	if err != nil {
		t.Error(err)
	}

	t.Log("album modified", chg, oneAlbum)

	oneAlbum, err = albumDao.Get(oneAlbum.Id)
	if err != nil {
		t.Error(err)
	}

	t.Log("album found one modified", oneAlbum)

	err = albumDao.Delete(oneAlbum.Id)
	if err != nil {
		t.Error(err)
	}

	oneAlbum, err = albumDao.Get(oneAlbum.Id)
	if err != nil {
		t.Log("album deleted", err, oneAlbum)
	}
}
