import { useState, useRef, useEffect } from "react";

export const WebSocketConsole = ({ port }) => {
  const [ws, setWs] = useState(null);
  const [incomingMessage, setIncomingMessage] = useState("");
  const [editableText, setEditableText] = useState("");
  const [errors, setErrors] = useState([]);
  const [loggedMessages, setLoggedMessages] = useState([]);
  const [syncConnection, setSyncConnection] = useState(true);
  let incomingMessageJSON = useRef({});

  // After the element rendering, connect to the given websocket port
  useEffect(() => {
    const newWs = new WebSocket(`ws://localhost:${port}/ws`);

    newWs.addEventListener("open", handleOpenWebSocket);
    newWs.addEventListener("message", handleMessageWebSocket);
    newWs.addEventListener("close", handleCloseWebSocket);
    newWs.addEventListener("error", handleErrorWebSocket);

    setWs(newWs);
    return () => {
      console.log("launching webSocket cleaning up part");
      if (ws) {
        ws.close();
        ws.removeEventListener("open", handleOpenWebSocket);
        ws.removeEventListener("message", handleMessageWebSocket);
        ws.removeEventListener("close", handleCloseWebSocket);
        ws.removeEventListener("error", handleErrorWebSocket);
      }
    };
  }, []);

  // When a new incomingMessage log is given, parse it to editable
  useEffect(() => {
    if (incomingMessage !== null && incomingMessage !== "") {
      incomingMessageJSON.current = JSON.parse(incomingMessage);
      setEditableText(
        jsonToEditableText(incomingMessageJSON.current, null, "", 0)
      );
    }
  }, [incomingMessage]);

  // Need a different useEffect, due to the different dependency set
  useEffect(() => {
    if (ws) {
      ws.addEventListener("message", handleMessageWebSocket);
    }

    return () => {
      if (ws) {
        ws.removeEventListener("message", handleMessageWebSocket);
      }
    };
  }, [syncConnection]);

  function handleOpenWebSocket() {
    const message = `WebSocket connection for Port ${port} established.`;
    addLoggedMessage(message);
    console.log(message);

    // Clear error messages
    setErrors([]);
  }

  function handleMessageWebSocket(event) {
    const message = decomposeJSON(event.data);
    if (syncConnection === true) {
      // Sync mode
      console.log("currently in sync and waiting", port);
      setIncomingMessage(message); // This will launch the useEffect dependent on IncomingMessage
    } else {
      // Async mode
      acceptAsyncIncomingLog(message);
    }
  }

  function handleErrorWebSocket(event) {
    const message = `WebSocket error. See console for further details`;
    console.error(event.valueOf());
    console.log(errors);
    setErrors((prevErrors) => [...prevErrors, message]);
  }

  function handleCloseWebSocket() {
    const message = `WebSocket connection ended`;
    setErrors((prevErrors) => [...prevErrors, message]);
    console.log(`WebSocket connection for Port ${port} closed.`);
  }

  const addLoggedMessage = (newMessage) => {
    setLoggedMessages((prevMessages) => [newMessage, ...prevMessages]);
  };

  const decomposeJSON = (message) => {
    // Note: This decompose function is closely tied to the composition in websocketwriter.go.
    // We expect the following structure:
    // {
    //      event: string(message),
    //      timestamp: "sending_timestamp"
    // }
    // Here, we need to perform double parsing because we need to re-deparse the 'message' as it was originally in another JSON format.
    let parsedMessage = JSON.parse(message);
    parsedMessage.event = JSON.parse(parsedMessage.event);
    return JSON.stringify(parsedMessage, null, 2);
  };

  const clearIncomingLogMessages = () => {
    setIncomingMessage("");
    setEditableText("");
    incomingMessageJSON = {};
  };

  const acceptIncomingLog = () => {
    if (incomingMessage === "") {
      return;
    }
    console.log("new message accepted on port: ", port);

    // Accept ORIGINAL Log (without modifications)
    addLoggedMessage(incomingMessage);

    // Send the response on the WebSocket
    let webSocketResponse = {
      Type: "accept",
      Value: "",
    };
    const webSocketResponseJSON = JSON.stringify(webSocketResponse);
    ws.send(webSocketResponseJSON);

    // Clear the input
    clearIncomingLogMessages();
  };

  const acceptAsyncIncomingLog = (message) => {
    console.log("new message async accepted on port: ", port);

    // Accept ORIGINAL Log (without modifications and automatically)
    addLoggedMessage(message);

    // Send the response on the WebSocket
    let webSocketResponse = {
      Type: "accept",
      Value: "",
    };
    const webSocketResponseJSON = JSON.stringify(webSocketResponse);
    ws.send(webSocketResponseJSON);

    // Clear the input
    clearIncomingLogMessages();
  };

  const replaceIncomingLog = () => {
    if (incomingMessage === "") {
      return;
    }

    // Accept REPLACED Log
    const acceptingModifiedMessage = JSON.stringify(
      incomingMessageJSON.current,
      null,
      2
    );
    addLoggedMessage(acceptingModifiedMessage);

    // Send the response on the WebSocket

    // Note: This parsing function is closely tied to the composition in websocketwriter.go.
    // We expect the following structure:
    // {
    //      event: string(message),
    //      timestamp: "sending_timestamp"
    // }
    // Here, we need to perform double parsing for the 'event' value as it was originally also in JSON format.
    incomingMessageJSON.current.event = JSON.stringify(
      incomingMessageJSON.current.event
    );
    let webSocketResponse = {
      Type: "replace",
      Value: JSON.stringify(incomingMessageJSON.current),
    };
    const webSocketResponseJSON = JSON.stringify(webSocketResponse);
    console.log(
      "webSocketResponseJSON",
      webSocketResponseJSON,
      webSocketResponse
    );
    ws.send(webSocketResponseJSON);

    // Clear the input
    clearIncomingLogMessages();
  };

  const declineIncomingLog = () => {
    if (incomingMessage === "") {
      return;
    }

    // Send the decline response on the WebSocket
    let webSocketResponse = {
      Type: "decline",
      Value: "",
    };
    const webSocketResponseJSON = JSON.stringify(webSocketResponse);
    ws.send(webSocketResponseJSON);

    // Clear the input
    clearIncomingLogMessages();
  };

  // Modifiable text component: When clicking on it, it becomes an input text
  const EditableText = ({
    text,
    textType,
    jsonReference,
    jsonReferenceKey,
    onUpdate,
  }) => {
    if (text === null || text === "undefined") {
      text = ""; // Set text to an empty string if it's null or undefined
    }
    const [isEditing, setIsEditing] = useState(false);
    const [editedText, setEditedText] = useState(text);
    const editedTextType = useRef(textType);
    const textRef = useRef();
    const myJsonReference = useRef(jsonReference);
    const myJsonReferenceKey = useRef(jsonReferenceKey);
    console.log(
      "Initializing myJsonReference",
      myJsonReference,
      text,
      jsonReference
    );

    const handleDoubleClick = () => {
      setIsEditing(true);
    };

    const handleBlur = () => {
      setIsEditing(false);

      // Update parent with the new text
      onUpdate(
        editedText,
        editedTextType.current,
        myJsonReference.current,
        myJsonReferenceKey.current
      );
    };

    const handleChange = (e) => {
      setEditedText(e.target.value);
    };

    return (
      <label onClick={handleDoubleClick}>
        {isEditing ? (
          <input
            type="text"
            value={editedText}
            onChange={handleChange}
            onBlur={handleBlur}
            ref={textRef}
          />
        ) : (
          <span className="editable-text-span">{editedText}</span>
        )}
      </label>
    );
  };

  // Recursive function to parse the JSON and convert the leafs into EditableText
  // Remember that the JSON should be passed by reference in JavaScript; therefore, we can modify the same object.
  // Note that if the text is a boolean, it will be converted firstly to a string. After modification,
  //      if it is still an acceptable value for a boolean, it will be converted back.
  // Note that if the text is a number, it will be converted firstly to a string. After modification,
  //      if it is still an acceptable value for a number, it will be converted back.
  // Note that if the text is null, once the text is clicked, it will become an empty string.
  const jsonToEditableText = (json, parentJson, parentKey, depth) => {
    const space = "  "; // Two spaces for each level of depth
    const indentation = space.repeat(depth);
    if (typeof json === "object" && json !== null) {
      if (Array.isArray(json)) {
        // Array object
        return (
          <div>
            {indentation}&#91;
            {json.map((item, index) => {
              return (
                <div key={index}>
                  {indentation}
                  {space}
                  {jsonToEditableText(item, json, index, depth + 1)}
                </div>
              );
            })}
            {indentation}&#93;
          </div>
        );
      } else if (Object.keys(json).length === 0) {
        // Empty object
        // Return inline {}
        return <span> &#123;&#125;</span>;
      } else {
        // Normal non-empty Object
        return (
          <div>
            {indentation}&#123;
            {Object.keys(json).map((key) => {
              return (
                <div key={key}>
                  {indentation}
                  <strong>
                    {space}"{key}":
                  </strong>{" "}
                  {jsonToEditableText(json[key], json, key, depth + 1)}
                </div>
              );
            })}
            {indentation}&#125;
          </div>
        );
      }
    } else {
      // Manage json type edge cases
      // Skipped the edge case of given a number, as a number will be automatically converted to string
      // Manage the case is given a boolean and convert it back to string to print
      const textType = typeof json;
      if (textType === "boolean") {
        json = json ? "true" : "false";
      }
      return (
        <EditableText
          text={json}
          textType={textType}
          jsonReference={parentJson}
          jsonReferenceKey={parentKey}
          onUpdate={(newText, textType, jsonReference, jsonReferenceKey) => {
            // On screen print it as string but on the background save it with the actual type
            json = newText; // String on screen
            let mappingValue = newText;
            if (textType === "boolean") {
              if (mappingValue === "true") {
                mappingValue = true;
              } else if (mappingValue === "false") {
                mappingValue = false;
              } else {
                // The boolean has been converted to a string
                textType = "string";
              }
            } else if (textType === "number") {
              const temp = Number(mappingValue);
              if (isNaN(temp)) {
                // Not a valid Number
                textType = "string";
              } else {
                mappingValue = temp;
              }
            }
            jsonReference[jsonReferenceKey] = mappingValue; // Actual value (NOT String)
          }}
        />
      );
    }
  };

  const closeConnection = () => {
    console.log("Closing connection from button");
    if (ws) {
      declineIncomingLog();
      ws.close();

      // Send the response on the WebSocket
      let webSocketResponse = {
        Type: "close",
        Value: "",
      };
      const webSocketResponseJSON = JSON.stringify(webSocketResponse);
      ws.send(webSocketResponseJSON);
    }
  };

  const changeSyncronization = () => {
    console.log("changing syncronization of port: ", port);
    if (syncConnection) {
      // Changing to Async mode
      acceptIncomingLog(); // Accept also the existing incoming log
    }
    setSyncConnection(!syncConnection);
  };

  return (
    <div id="sub-screen" className="sub-screen">
      <label className="switch">
        <input
          type="checkbox"
          id="sync-switch-input"
          checked={syncConnection}
          onChange={changeSyncronization}
        />
        <div className="slider round">
          <span className="on">Sync</span>
          <span className="off">Async</span>
        </div>
      </label>
      <button onClick={closeConnection} className="close-button">
        Close Connection
      </button>
      <p className="websocket-console-title">
        WebSocket Console for Port {port}:
      </p>

      <div className="error-screen">
        {errors.map((error, index) => (
          <p key={index}>{error}</p>
        ))}
      </div>

      <div id="incoming-log-div" className="incoming-log-div">
        Incoming Log:
        <div
          id="incoming-log-placeholder"
          className="incoming-log-screen json-message"
        >
          <pre>{editableText}</pre>
        </div>
        <div id="incoming-log-buttons" className="incoming-log-buttons">
          <button className="button" onClick={acceptIncomingLog}>Accept</button>
          <button className="button" onClick={replaceIncomingLog}>Replace</button>
          <button className="button" onClick={declineIncomingLog}>Decline</button>
        </div>
      </div>

      <div id="logged-messages-div" className="logged-messages-div">
        Accepted Logs:
        <div id="log-placeholder" className="log-screen json-message">
          {loggedMessages.map((log, index) => (
            <p key={index} className="log-message">
              {log}
            </p>
          ))}
        </div>
      </div>
    </div>
  );
};
