import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, catchError, of } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class HttpHandlerService {
  baseUrl: string = 'http://localhost:8080/' 
  constructor(private http: HttpClient) {}

  get<T>(url: string, options?: any, err: string = ''): Observable<any> {
    return this.http.get<T>(url, options).pipe(catchError(this.handleError(err)));
  }

  post<T>(url: string, body: any, options?: any, err: string = ''): Observable<any> {
    return this.http.post<T>(url, body, options).pipe(catchError(this.handleError(err)));
  }

  private handleError<T> (request: string, result?: T): (error: Error) => Observable<T> {
    return () => {
      return of(result as T);
    };
  }
}
