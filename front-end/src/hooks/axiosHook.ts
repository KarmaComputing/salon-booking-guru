// vue
import { useStore } from 'vuex';

// axios
import axios, { AxiosError } from 'axios';

// import { useToast } from 'vue-toastification';

export const useAxios = () => {
    // hooks
    const store = useStore();
    // const toast = useToast();

    // methods
    const handleError = (err: AxiosError) => {
        if (err.response && err.response.data.message) {
            // toast.error(err.response.data.message);
        } else if (err.response && err.response.data.errors) {
            let errors = '';
            err.response.data.errors.forEach((error: string, i: number) => {
                errors += '    ' + error;
                if (i !== errors.length - 1) {
                    errors += '\n';
                }
            });
            // toast.error('Validation error:\n' + errors);
        } else {
            // toast.error(JSON.stringify(err));
        }
        throw err;
    };

    // eslint-disable-next-line
    const get = (url: string): any => {
        const authToken = store?.state?.authToken;
        if (!authToken) {
            return null;
        }

        const axiosInstance = axios
            .get(url, {
                cancelToken: store.state.cancelToken.token,
                headers: {
                    Authorization: 'Bearer ' + JSON.parse(authToken).token,
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    // eslint-disable-next-line
    const post = (url: string, data: any): any => {
        const authToken = store.getters.authToken;

        const axiosInstance = axios
            .post(url, data, {
                cancelToken: store.state.cancelToken.token,
                headers: {
                    Authorization: 'Bearer ' + authToken?.token,
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    // eslint-disable-next-line
    const postMultipart = (url: string, data: any): any => {
        const authToken = store?.state?.authToken;
        if (!authToken) {
            return null;
        }

        const axiosInstance = axios
            .post(url, data, {
                cancelToken: store.state.cancelToken.token,
                headers: {
                    Authorization: 'Bearer ' + JSON.parse(authToken).token,
                    'Content-Type': 'multipart/form-data',
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    // eslint-disable-next-line
    const put = (url: string, data: any): any => {
        const authToken = store?.state?.authToken;
        if (!authToken) {
            return null;
        }

        const axiosInstance = axios
            .put(url, data, {
                cancelToken: store.state.cancelToken.token,
                headers: {
                    Authorization: 'Bearer ' + JSON.parse(authToken).token,
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    // eslint-disable-next-line
    const deleteId = (url: string): any => {
        const authToken = store?.state?.authToken;
        if (!authToken) {
            return null;
        }

        const axiosInstance = axios
            .delete(url, {
                cancelToken: store.state.cancelToken.token,
                headers: {
                    Authorization: 'Bearer ' + JSON.parse(authToken).token,
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    // eslint-disable-next-line
    const deleteObject = (url: string, data: any): any => {
        const authToken = store?.state?.authToken;
        if (!authToken) {
            return null;
        }

        const axiosInstance = axios
            .delete(url, {
                cancelToken: store.state.cancelToken.token,
                data,
                headers: {
                    Authorization: 'Bearer ' + JSON.parse(authToken).token,
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    // eslint-disable-next-line
    const download = (url: string): any => {
        const authToken = store?.state?.authToken;
        if (!authToken) {
            return null;
        }

        const axiosInstance = axios
            .get(url, {
                cancelToken: store.state.cancelToken.token,
                responseType: 'blob',
                headers: {
                    Authorization: 'Bearer ' + JSON.parse(authToken).token,
                },
            })
            .catch((err: AxiosError) => {
                handleError(err);
            });
        return axiosInstance;
    };

    return {
        get,
        post,
        postMultipart,
        put,
        deleteObject,
        deleteId,
        download,
    };
};
