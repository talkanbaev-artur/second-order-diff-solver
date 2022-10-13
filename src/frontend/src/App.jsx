import { useEffect, useState } from "react";
import api from "./api";
import { LinePlot, createDataSet } from "./components/chart";


function App() {

  const [backendData, setBackendData] = useState(createDataSet([], [], "original", "red"));
  const [numerical, setNumerical] = useState(createDataSet([], [], "numerical", "green"));

  useEffect(() => {
    const func = async () => {
      var data = await api.getAnalytical("3", 1);
      var dataNum = await api.getNumerical("3", 1, 1024, "central")
      if (data.status == 200 && dataNum.status == 200) {
        setBackendData(createDataSet(data.data.xVals, data.data.yVals, "original", "red"));
        setNumerical(createDataSet(dataNum.data.xVals, dataNum.data.yVals, "numerical", "green"))
      };
    }
    func()
    return () => { }
  }, []);

  return (
    <div className="">
      <div className="">
        {LinePlot([backendData, numerical])}
      </div>
    </div>
  )
}

export default App
