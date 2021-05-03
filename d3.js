d3.json("/get_timeline_data", {
    method: "POST",
    body: JSON.stringify({
        date: Date.now(),
    })
}).then((data) => {
    console.log(data)
}).catch((error) => {
    console.error("Error loading the data: " + error);
});
