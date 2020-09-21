package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/restapiGo/mahasiswa"
	"github.com/restapiGo/models"
	"github.com/restapiGo/utils"
)

// GetMahasiswa
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "d!am$y@h-m-d1d4-k3cE-p4R@h!!!" {
		if r.Method == "GET" {
			ctx, cancel := context.WithCancel(context.Background())

			defer cancel()

			mahasiswas, err := mahasiswa.GetAll(ctx)

			if err != nil {
				fmt.Println(err)
			}

			utils.ResponseJSON(w, mahasiswas, http.StatusOK)
			return
		}

		http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
		return

	} else {
		http.Error(w, "key invalid", http.StatusUnauthorized)
		return
	}
}

// PostMahasiswa
func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "d!am$y@h-m-d1d4-k3cE-p4R@h!!!" {
		if r.Method == "POST" {

			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
				return
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var mhs models.Mahasiswa

			if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
				utils.ResponseJSON(w, err, http.StatusBadRequest)
				return
			}

			if err := mahasiswa.Insert(ctx, mhs); err != nil {
				utils.ResponseJSON(w, err, http.StatusInternalServerError)
				return
			}

			res := map[string]string{
				"status": "Succesfully",
			}

			utils.ResponseJSON(w, res, http.StatusCreated)
			return
		}

		http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
		return

	} else {
		http.Error(w, "key invalid", http.StatusUnauthorized)
		return
	}
}

// UpdateMahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "d!am$y@h-m-d1d4-k3cE-p4R@h!!!" {
		if r.Method == "PUT" {

			if r.Header.Get("Content-Type") != "application/json" {
				http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
				return
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var mhs models.Mahasiswa

			if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
				utils.ResponseJSON(w, err, http.StatusBadRequest)
				return
			}

			fmt.Println(mhs)

			if err := mahasiswa.Update(ctx, mhs); err != nil {
				utils.ResponseJSON(w, err, http.StatusInternalServerError)
				return
			}

			res := map[string]string{
				"status": "Succesfully",
			}

			utils.ResponseJSON(w, res, http.StatusCreated)
			return
		}

		http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
		return

	} else {
		http.Error(w, "key invalid", http.StatusUnauthorized)
		return
	}
}

// DeleteMahasisw
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "d!am$y@h-m-d1d4-k3cE-p4R@h!!!" {
		if r.Method == "DELETE" {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var mhs models.Mahasiswa

			id := r.URL.Query().Get("id")

			if id == "" {
				utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
				return
			}
			mhs.ID, _ = strconv.Atoi(id)

			if err := mahasiswa.Delete(ctx, mhs); err != nil {

				kesalahan := map[string]string{
					"error": fmt.Sprintf("%v", err),
				}

				utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
				return
			}

			res := map[string]string{
				"status": "Succesfully",
			}

			utils.ResponseJSON(w, res, http.StatusOK)
			return
		}

		http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
		return

	} else {
		http.Error(w, "key invalid", http.StatusUnauthorized)
		return
	}
}
