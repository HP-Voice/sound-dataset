const API_PORT = "4431"
export const API_URL = window.location.protocol + "//" + window.location.hostname + ":" + API_PORT + "/"
export function getRandomSpell() {
    return new Promise(
        (resolve, reject) => {
            fetch(API_URL+"random-spell")
                .then(async response => {
                    if (response.status !== 200) {
                        reject(response.status + " " + await response.text());
                    }
                    resolve(JSON.parse(await response.text()));
                })
                .catch(reason => reject(reason));
        }
    );
}
export function getLabels() {
    return new Promise(
        (resolve, reject) => {
            fetch(API_URL+"labels")
                .then(async response => {
                    if (response.status !== 200) {
                        reject(response.status + " " + await response.text());
                    }
                    resolve(JSON.parse(await response.text()));
                })
                .catch(reason => reject(reason));
        }
    );
}
export function postSample(labelId, data) {
    let formData = new FormData();
    formData.set("Label-Id", labelId);
    formData.set("Sample", data, "sample.ogg");
    return new Promise(
        (resolve, reject) => {
            fetch(API_URL+"sample", {
                method: "POST",
                body: formData,
            })
                .then(async response => {
                    if (response.status !== 200) {
                        reject(response.status + " " + await response.text());
                    }
                    resolve();
                })
                .catch(reason => reject(reason));
        }
    );
}
export function getSentence() {
    return new Promise(
        (resolve, reject) => {
            fetch(API_URL + "sentence")
                .then(async response => {
                    if (response.status !== 200) {
                        reject(response.status + " " + await response.text());
                    }
                    resolve(await response.text());
                })
                .catch(reason => reject(reason))
        }
    )
}