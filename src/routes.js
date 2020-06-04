const express = require('express')

const weatherController = require('./controllers/appController')

const routes = express.Router()

routes.get('/playlist', weatherController.getByCityName)

module.exports = routes
