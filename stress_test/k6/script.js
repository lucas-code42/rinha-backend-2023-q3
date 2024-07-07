import http from "k6/http";
import { sleep, check } from "k6";
import { randomString, randomItem } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export let options = {
    stages: [
        { duration: "05s", target: 1 },
        { duration: "10s", target: 10 },
        { duration: "20s", target: 500 },
        { duration: "40s", target: 50 },
    ],
};

function generateRandomDate(startYear, endYear) {
    const year = Math.floor(Math.random() * (endYear - startYear + 1)) + startYear;
    const month = ("0" + (Math.floor(Math.random() * 12) + 1)).slice(-2);
    const day = ("0" + (Math.floor(Math.random() * 28) + 1)).slice(-2);
    return `${year}/${month}/${day}`;
}

function generateRandomStack() {
    const technologies = [
        "golang", "python", "rust", "java", "javascript", "c++", "c", "elixir", "c#", "ruby", "kotlin"
    ];
    let stack = [];
    for (let i = 0; i < 3; i++) {
        stack.push(randomItem(technologies));
    }
    return [...new Set(stack)];
}

function generateRandomSearchTerm() {
    const length = Math.floor(Math.random() * 10) + 1;
    return randomString(length);
}

export default function () {
    const apelido = randomString(10);
    const nome = randomString(10) + " " + randomString(4);
    const stack = generateRandomStack();
    const nascimento = generateRandomDate(1945, 1989);

    let data = JSON.stringify({
        "apelido": apelido,
        "nome": nome,
        "stack": stack,
        "nascimento": nascimento
    })
    let createPersonRes = http.post(
        "http://localhost:9999/pessoas",
        data,
        {
            headers: { "Content-Type": "application/json" }
        }
    );

    console.log("DEBUG:", data);
    check(createPersonRes, { "POST /pessoas status was 201": (r) => r.status === 201 });

    let location = createPersonRes.headers["Location"];
    let personId = location ? location.split("/").pop() : null;
    if (personId) {
        let getPersonRes = http.get(`http://localhost:9999/pessoas/${personId}`);
        check(getPersonRes, { "GET /pessoas/{id} status was 200": (r) => r.status === 200 });
    }

    let searchTerm = generateRandomSearchTerm();
    let searchPersonRes = http.get(`http://localhost:9999/pessoas?t=${searchTerm}`);
    check(searchPersonRes, { "GET /pessoas?t={termo} status was 200": (r) => r.status === 200 });

    let countPersonRes = http.get("http://localhost:9999/contagem-pessoas");
    check(countPersonRes, { "GET /contagem-pessoas status was 200": (r) => r.status === 200 });

    sleep(1);
}
