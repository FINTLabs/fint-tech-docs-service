import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import { RepoModel } from './RepoModel';

@Injectable()
export class RepositoriesService {
  url = '/api/projects';
  isLoading = false;
  cache: RepoModel[];

  constructor(private http: Http) { }

  /**
   * This will only read the model once, and forever after
   * return the cached model. If you want a fresh model,
   * you will have to refresh your browser!
   */
  all(): Observable<RepoModel[]> {
    this.isLoading = true;
    if (this.cache) { return Observable.of(this.cache); }
    return this.http.get(`${this.url}`)
    // return this.http.get(`https://docs.felleskomponent.no${this.url}`)
      .map(result => {
        this.isLoading = false;
        this.cache = result.json();
        return this.cache;
      })
      .catch(error => {console.error(error); return error; });
  }

}
