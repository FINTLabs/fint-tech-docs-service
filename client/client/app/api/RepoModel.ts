export interface RepoModel {
  description: string;
  dirty: boolean;
  git: string;
  lang: 'java' | 'net';
  html: string;
  latest: string;
  latest_time: string;
  latest_url: string;
  maven: string;
  maven_badge: string;
  name: string;
  readme: string;
}
