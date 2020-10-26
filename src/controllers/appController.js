const config = require('../config/config')
const service = require('../services/appService')

const apiKey = config.openWeatherApiKey

module.exports={
    async getByCityName(req, res){
        const { city } = req.query

        let temp = await service.getPlaylist(city, apiKey)

        return res.json({temp})
    }
}
