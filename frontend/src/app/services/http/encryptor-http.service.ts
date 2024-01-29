import { Injectable } from '@angular/core';
import { HttpHandlerService } from './http-handler.service';
import { Encryptor } from 'src/app/models/interfaces/encryptor';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class EncryptorHttpService {

  constructor(private http: HttpHandlerService) { }

  encrypt(encryptObj: Encryptor): Observable<any> {
    return this.http.post<any>('encrypt', encryptObj.toDto().obj)
  }
}
