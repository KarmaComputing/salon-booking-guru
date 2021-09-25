// vue
import { createStore } from 'vuex';

// axios
import axios from 'axios';

export default createStore({
    state: {
        apiUrl: 'http://localhost:8085/v1',
        authToken: localStorage.getItem('authToken'),
        accountInfo: localStorage.getItem('accountInfo'),
        cancelToken: axios.CancelToken.source(),
    },
    getters: {
        authToken: (state) => {
            if (state.authToken) {
                return JSON.parse(state.authToken);
            }
        },
        accountInfo: (state) => {
            if (state.accountInfo) {
                return JSON.parse(state.accountInfo);
            }
        },
    },
    mutations: {
        accountInfo(state, value) {
            state.accountInfo = JSON.stringify(value);
            if (value) {
                localStorage.setItem('accountInfo', JSON.stringify(value));
            } else {
                localStorage.removeItem('accountInfo');
            }
        },
        authToken(state, value) {
            state.authToken = JSON.stringify(value);
            if (value) {
                localStorage.setItem('authToken', JSON.stringify(value));
            } else {
                localStorage.removeItem('authToken');
            }
        },
    },
    actions: {},
    modules: {},
});
