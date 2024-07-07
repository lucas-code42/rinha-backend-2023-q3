import http from "k6/http";
import { sleep, check } from "k6";
import { randomString, randomItem } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export let options = {
    stages: [
        { duration: "10s", target: 50 },
        { duration: "10s", target: 100 },
        { duration: "10s", target: 500 },
        { duration: "20s", target: 1000 },
        { duration: "10s", target: 10 },
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

    console.log(data);
    check(createPersonRes, { "POST /pessoas status was 201": (r) => r.status === 201 });

    sleep(1);
}
