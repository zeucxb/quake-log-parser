package report

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quake-log-parser/helper"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var jsonFilePath = "./quake_data.json"

func reportDefaultHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	message := "API ONLINE\n\nAcesse /report para obter o relatório completo dos logs\nAcesse /report/:gameName ou /report/:gameNumber para obter o relatório do log de um game"

	fmt.Fprintf(w, message)
}

func reportAllHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	games, err := helper.GetGames(jsonFilePath)
	if checkError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func reportByGameHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	games, err := helper.GetGames(jsonFilePath)
	if checkError(w, err) {
		return
	}

	gameName := sanitalizeGameName(ps.ByName("game"))

	if game, ok := games[gameName]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// Serve a the report api
func Serve(portStr string) (err error) {
	router := httprouter.New()
	router.GET("/", reportDefaultHandle)
	router.GET("/report", reportAllHandle)
	router.GET("/report/:game", reportByGameHandle)

	portStr = fmt.Sprintf(":%v", portStr)

	err = http.ListenAndServe(portStr, router)
	return
}

func checkError(w http.ResponseWriter, err error) (ok bool) {
	ok = helper.CheckError(err)
	if ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	return
}

func sanitalizeGameName(gameName string) (sanitalizedGameName string) {
	if !strings.HasPrefix(gameName, "game_") {
		return fmt.Sprintf("game_%s", gameName)
	}

	return gameName
}
