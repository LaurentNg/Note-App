import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, of, throwError } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class HttpHandlerService {
  private baseUrl: string = 'http://localhost:8080/' 
  constructor(private http: HttpClient) {}

  get<T>(url: string, options?: any, err: string = ''): Observable<any> {
    return this.http.get<T>(this.baseUrl + url, options).pipe(catchError(this.handleError));
  }

  post<T>(url: string, body: any, options?: any, err: string = ''): Observable<any> {
    return this.http.post<T>(this.baseUrl + url, body, options).pipe(catchError(this.handleError));
  }

  // private handleError<T> (request: string, result?: T): (error: Error) => Observable<T> {
  //   return () => {
  //     return of(result as T);
  //   };
  // }

  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return of(error.error);
  }
}
