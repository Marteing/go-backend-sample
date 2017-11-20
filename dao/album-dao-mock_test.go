package dao_test

import (
	"go-backend-sample/dao"
	"go-backend-sample/model"
	"testing"
)

func TestAlbumDAOMock(t *testing.T) {

	_, albumDaoMock, err := dao.GetDAO(dao.MockDAO, "")
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
		Id:          "1",
		Title:       "Test",
		Description: "Description Test",
		AuthorId:    "1",
		Songs:       songsToSave,
	}

	albumSaved, err := albumDaoMock.Upsert(&albumToSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("album saved", albumSaved)

	albums, err := albumDaoMock.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(albums) != 1 {
		t.Errorf("expected 1 albums, got %d", len(albums))
	}

	oneAlbum, err := albumDaoMock.Get(albumToSave.Id)
	if err != nil {
		t.Error(err)
	}
	if albumSaved != oneAlbum {
		t.Error("got wrong album by id")
	}

	err = albumDaoMock.Delete(oneAlbum.Id)
	if err != nil {
		t.Error(err)
	}

	oneAlbum, err = albumDaoMock.Get(oneAlbum.Id)
	if err == nil {
		t.Error("album should have been deleted", oneAlbum)
	}
}
