const axios = require('axios')
const tuc = require('temp-units-conv')

// function selectPlaylist(temp){
//     if(temp > 30){

//     }else if(temp < 30 && temp > 15){

//     }else if(temp < 15 && temp > 10){

//     }else{

//     }
// }

module.exports = {
    async getPlaylist(city, apiKey){
        try{
            const res = await axios.get(`http://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${apiKey}`)

            const temp = res.data.main['temp']
            temp = tuc.k2c(temp)


        }catch(error){
            console.log(error)
        }
    }
}
