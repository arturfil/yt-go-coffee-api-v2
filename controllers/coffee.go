package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arturfil/yt-go-coffee-server/helpers"
	"github.com/arturfil/yt-go-coffee-server/services"
	"github.com/go-chi/chi/v5"
)


var coffee services.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
    all, err := coffee.GetAllCoffees()
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

// GET//coffees/coffee/{id}
func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    coffee, err := coffee.GetCoffeeById(id)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, coffee)
}

// POST/coffees/coffee
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
    var coffeeData services.Coffee
    err := json.NewDecoder(r.Body).Decode(&coffeeData)

    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return 
    }

    coffeeCreated, err := coffee.CreateCoffee(coffeeData)
    // CHECK
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}

// PUT/coffees/coffee/{id}
func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
    var coffeeData services.Coffee
    id := chi.URLParam(r, "id")
    err := json.NewDecoder(r.Body).Decode(&coffeeData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    coffeUpdated, err := coffee.UpdateCoffee(id, coffeeData)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
    }
    helpers.WriteJSON(w, http.StatusOK, coffeUpdated)
}

// DELETE/coffees/coffee/{id}
func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    err := coffee.DeleteCoffee(id)
    if err != nil {
       helpers.MessageLogs.ErrorLog.Println(err) 
    }
    helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}
