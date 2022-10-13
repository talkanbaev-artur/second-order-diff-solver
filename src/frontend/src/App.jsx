import { useEffect, useState } from "react";
import api from "./api";
import { LinePlot, createDataSet } from "./components/chart";


function App() {

  const [backendData, setBackendData] = useState(createDataSet([], [], "original", "red"));

  useEffect(() => {
    const func = async () => {
      var data = await api.getAnalytical("3", 0.5);
      if (data.status == 200) {
        var original = createDataSet(data.data.xVals, data.data.yVals, "original", "rgb(255,10,10)")
        setBackendData(original);
      };
    }
    func()
    return () => { }
  }, []);

  return (
    <div className="">
      <div className="">
        {LinePlot([backendData])}
      </div>
    </div>
  )
}

export default App
