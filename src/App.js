import './App.css';
import Webcam from "react-webcam";
import React, {useEffect, useRef} from "react";
import {connect, sendMsg} from "./conn";

function App() {
    connect()

    const webcamRef = useRef(null);
    const canvasRef = useRef(null);

    // Main function
    const capture = () => {
        console.log("Handpose model loaded.");
        //  Loop and detect hands
        setInterval(() => {
            capFrame();
        }, 100);
    };

    const capFrame = () => {
        // Check data is available
        if (
            typeof webcamRef.current !== "undefined" &&
            webcamRef.current !== null &&
            webcamRef.current.video.readyState === 4
        ) {
            const pic = webcamRef.current.getScreenshot();
            // console.log("data is: ", pic)
            sendMsg(pic)
        }
    };

    useEffect(()=>{capture()});

  return (
      <div className="App">
        <header className="App-header">
          <Webcam
              ref={webcamRef}
              muted={true}
              audio={false}
              screenshotFormat={"image/jpeg"}
              style={{
                position: "absolute",
                marginLeft: "auto",
                marginRight: "auto",
                left: 0,
                right: 0,
                textAlign: "center",
                zindex: 9,
                width: 640,
                height: 480,
              }}
          />

          <canvas
              ref={canvasRef}
              style={{
                position: "absolute",
                marginLeft: "auto",
                marginRight: "auto",
                left: 0,
                right: 0,
                textAlign: "center",
                zindex: 8,
                width: 640,
                height: 480,
              }}
          />
        </header>
      </div>
  );
}

export default App;
