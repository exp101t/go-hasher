addEventListener("DOMContentLoaded", function () {
    let $ = (x) => document.querySelector(x);

    $('#submit_button').onclick = function () {
        let data = $('#data').value;
        let hashType = $('#hash_type').value;

        fetch(`/api/hash/${hashType}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({data: data})
        }).then((response) => response.json())
          .then((responseData) => {
            if (responseData.error) {
                $('#hash').innerText = `Error: ${responseData.error}`;
            } else {
                $('#hash').innerText = `Hash: ${responseData.hash}`;
            }
        })
    }
});
