chrome.runtime.sendMessage({ type: "GET_LOCALSTORAGE_DATA" }, (response) => {
  if (response && response.data) {
    for (const key in response.data) {
      if (Object.prototype.hasOwnProperty.call(response.data, key)) {
        localStorage.setItem(key, response.data[key]);
      }
    }
    console.log("[Sync BRI to Localhost] localStorage injected by content script.");
  } else {
    console.warn("[Sync BRI to Localhost] No localStorage data received.");
  }
});
