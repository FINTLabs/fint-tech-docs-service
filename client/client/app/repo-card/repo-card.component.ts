import { Component, OnInit, Input } from '@angular/core';
import { RepoModel } from '../api/RepoModel';

/**
 * Component responsible for displaying the details of
 * a repository. It expects to receive a `RepoModel` object
 * as input:
 *
 * ```html
 * <app-repo-card [repo]="repoModel"></app-repo-card>
 * ```
 */
@Component({
  selector: 'app-repo-card',
  templateUrl: './repo-card.component.html',
  styleUrls: ['./repo-card.component.scss']
})
export class RepoCardComponent implements OnInit {
  @Input() repo: RepoModel;

  constructor() { }

  ngOnInit() {
  }

}
