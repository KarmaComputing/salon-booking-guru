import { useAccountService } from './accountService';
import { useAuthenticateService } from './authenticateService';

export const useService = () => {
    return {
        ...useAccountService(),
        ...useAuthenticateService(),
    };
};
