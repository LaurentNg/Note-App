import { TestBed } from '@angular/core/testing';

import { EncryptorHttpService } from './encryptor-http.service';

describe('EncryptorHttpService', () => {
  let service: EncryptorHttpService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(EncryptorHttpService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
