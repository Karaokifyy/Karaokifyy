import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpInterceptor } from '@angular/common/http';
import {HttpEvent, HttpHandler, HttpRequest} from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable()
export class NoopInterceptorService implements HttpInterceptor {
    intercept(req: HttpRequest<any>, next: HttpHandler):
    Observable<HttpEvent<any>> {
    console.log(next.handle(req));
    //req.params.getAll.toString
    return next.handle(req);
  }
}
//console.log(req.params.getAll.toString);