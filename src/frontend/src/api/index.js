const apiEndpoint = "http://localhost:8000" //leave empty for prod

import axios from "axios"


const getTypes = () => {
    return axios.get(apiEndpoint+"/numericals")
}

export default {getTypes}