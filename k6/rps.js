import { sleep } from 'k6';
import http from 'k6/http';

export const options = {
    scenarios: {
        my_scenario1: {
            executor: 'constant-arrival-rate',
            duration: '13s', // total duration
            preAllocatedVUs: 250, // to allocate runtime resources  
            rate: 9000, // number of constant iterations given `timeUnit`
            timeUnit: '1s',
        },
    },
};

let url = 'http://localhost:5000/api/v1/ad';

export default function () {
    http.get(url)
}
