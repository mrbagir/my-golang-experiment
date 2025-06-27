const processedDomains = new Map();
const COOLDOWN_PERIOD_MS = 5000;
let latestLocalStorageData = {};

function getLocalhostDomain(briDomain) {
  if (briDomain.includes('.bri.co.id')) {
    return briDomain.replace('.bri.co.id', '.localhost');
  }
  return null;
}

async function copyLocalStorage(briDomain) {
  const lastProcessedTime = processedDomains.get(briDomain);
  if (lastProcessedTime && (Date.now() - lastProcessedTime < COOLDOWN_PERIOD_MS)) {
    console.log(`[Sync BRI to Localhost] Skipping ${briDomain}. Cooldown active.`);
    return;
  }

  try {
    const [briTab] = await chrome.tabs.query({ url: `*://${briDomain}/*` });
    if (!briTab) {
      console.warn(`[Sync BRI to Localhost] No tab found for ${briDomain}`);
      return;
    }

    const [briResult] = await chrome.scripting.executeScript({
      target: { tabId: briTab.id },
      func: () => JSON.stringify(localStorage)
    });

    latestLocalStorageData = JSON.parse(briResult.result);
    console.log("[Sync BRI to Localhost] localStorage cached:", latestLocalStorageData);

    processedDomains.set(briDomain, Date.now());
  } catch (error) {
    console.error(`[Sync BRI to Localhost] Error while copying from ${briDomain}:`, error);
  }
}

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message.type === "GET_LOCALSTORAGE_DATA") {
    sendResponse({ data: latestLocalStorageData });
  }
});

chrome.tabs.onActivated.addListener(async (activeInfo) => {
  const tab = await chrome.tabs.get(activeInfo.tabId);
  if (tab.url) {
    try {
      const url = new URL(tab.url);
      const hostname = url.hostname;
      if (hostname.includes('.bri.co.id')) {
        await copyLocalStorage(hostname);
      }
    } catch (e) {
      console.warn("[Sync BRI to Localhost] Failed to parse tab URL:", e);
    }
  }
});

chrome.tabs.onUpdated.addListener(async (tabId, changeInfo, tab) => {
  if (changeInfo.status === 'complete' && tab.url) {
    try {
      const url = new URL(tab.url);
      const hostname = url.hostname;
      if (hostname.includes('.bri.co.id')) {
        await copyLocalStorage(hostname);
      }
    } catch (e) {
      console.warn("[Sync BRI to Localhost] Failed to parse updated tab URL:", e);
    }
  }
});

console.log("[Sync BRI to Localhost] Background loaded.");
