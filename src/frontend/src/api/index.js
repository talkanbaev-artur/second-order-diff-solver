const apiEndpoint = "http://localhost:8000" //leave empty for prod

import axios from "axios"


const getTypes = () => {
    return axios.get(apiEndpoint + "/numericals")
}

const getAnalytical = (task, eps) => {
    return axios.post(apiEndpoint + "/analytical", { task, eps })
}

export default { getTypes, getAnalytical }