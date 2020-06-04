const config = require('../config/config')
const service = require('../services/appService')

const apiKey = config.openWeatherApiKey

module.exports={
    async getByCityName(req, res){
        console.log(apiKey)
        const { q } = req.body

        let temp = await service.getPlaylist(q, apiKey)

        return res.json({temp})
    }
}
