// models
import { AccountInfo } from './accountInfo';
import { Token } from './token';

export interface AuthenticateResponse {
    accountInfo: AccountInfo;
    token: Token;
}
