const apiEndpoint = "/api" //leave empty for prod

import axios from "axios"


const getTypes = () => {
    return axios.get(apiEndpoint + "/numericals")
}

const getAnalytical = (task, eps) => {
    return axios.post(apiEndpoint + "/analytical", { task, eps })
}

const getNumerical = (task, eps, n, scheme) => {
    return axios.post(apiEndpoint + "/solve", { task, eps, n, scheme })
}

export default { getTypes, getAnalytical, getNumerical }