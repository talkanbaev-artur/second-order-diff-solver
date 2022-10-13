import { Line } from "react-chartjs-2"
import { Chart, registerables } from 'chart.js';
import zoomPlugin from 'chartjs-plugin-zoom';
Chart.register(...registerables);
Chart.register(zoomPlugin);

//creates a data set
function createDataSet(xVals, yVals, label, color) {
    var dataset = { label: label, borderColor: color }
    if (xVals.length != yVals.length) {
        throw "invalid xy vals supplied: lengths do not match"
    }
    var data = []
    for (let i = 0; i < xVals.length; i++) {
        data[i] = { x: xVals[i], y: yVals[i] }
    }
    dataset.data = data;
    return dataset;
}

//supply an array of sets created by createDataSet function
function LinePlot(datasets) {
    var options = {
        animation: false,
        elements: {
            point: {
                radius: 0
            },
            line: {
                cubicInterpolationMode: "monotonic",
            }
        },
        responsive: true,
        scales: {
            x: {
                type: 'linear'
            },
            y: {
                type: 'linear',
                min: 0,
                max: 1
            }
        },
        plugins: {
            zoom: {
                zoom: {
                    wheel: {
                        enabled: true,
                    },
                    mode: 'xy',
                },
                pan: {
                    enabled: true,
                    mode: "xy",
                }
            }
        }
    }
    var data = {
        datasets: datasets
    }
    return (
        <Line data={data} options={options}></Line>
    )

}

export { LinePlot, createDataSet };