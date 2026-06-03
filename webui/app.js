const API = './api/v1';

const translations = {
  en: {
    sidebarNavAria: 'Main navigation',
    pageNavAria: 'Pages',
    languageToggleAria: 'Switch language',
    languageToggle: '中文',
    navDashboard: 'Dashboard',
    navLogs: 'Query Logs',
    navUpstreams: 'Upstreams',
    navConfig: 'Config',
    storageLabel: 'Storage',
    storageMemoryOnly: 'Memory only',
    storageRedisError: 'Redis error: {queue} queued',
    storageRedisActive: 'Redis active: {queue} queued',
    eyebrow: 'DNS observability',
    liveStatus: 'Live',
    refreshButton: 'Refresh',
    logoutButton: 'Logout',
    loginHeroKicker: 'Secure DNS Control',
    loginHeroText: 'View realtime queries, upstream status, and configuration after signing in.',
    loginEyebrow: 'WebUI access',
    loginTitle: 'Sign in',
    loginCopy: 'Use the WebUI username and password configured in mosdns.',
    loginUsername: 'Username',
    loginPassword: 'Password',
    loginButton: 'Sign in',
    loginWorking: 'Signing in...',
    loginFailed: 'Invalid username or password.',
    loginRequired: 'Please sign in to continue.',
    viewDashboard: 'Dashboard',
    viewLogs: 'Query Logs',
    viewUpstreams: 'Upstreams',
    viewConfig: 'Config Editor',
    metricTotalQueries: 'Total Queries',
    metricTotalQueriesHint: 'Processed queries',
    metricQps: 'QPS',
    metricQpsHint: '10-second average',
    metricAvgLatency: 'Avg Latency',
    metricAvgLatencyHint: 'End-to-end query time',
    metricCacheHitRate: 'Cache Hit Rate',
    cacheSizeHint: '{count} records',
    refreshCacheTitle: 'Refresh Domain Cache',
    refreshCacheDesc: 'Remove a domain from cache and query it again immediately.',
    cacheRefreshTypeAria: 'Cache refresh type',
    cacheRefreshEntryAria: 'Refresh entry plugin',
    refreshCacheButton: 'Refresh Cache',
    refreshCacheInitial: 'Enter a domain, then click refresh.',
    refreshCacheWorking: 'Refreshing cache...',
    refreshCacheSuccess: 'Cache refreshed',
    refreshCacheRequired: 'Domain is required',
    refreshRemovedSummary: '{count} cache key(s) removed · entry {entry}',
    noRefreshResult: 'No refresh result',
    hourlyQueriesTitle: 'Hourly Queries',
    hourlyQueriesDesc: 'Hourly query distribution',
    hourlyChartAria: 'Hourly query count',
    noHourlyData: 'No hourly query data yet',
    domainTopTitle: 'Domain Top 10',
    domainTopDesc: 'Most requested domains',
    noDomainData: 'No domain data yet',
    topClientsTitle: 'Top Clients',
    topClientsDesc: 'Client request ranking',
    noClientData: 'No client data yet',
    realtimeTailTitle: 'Realtime Query Tail',
    realtimeTailDesc: 'Latest 10 queries',
    openLogs: 'Open logs',
    queryLogsTitle: 'Realtime Query Logs',
    queryLogsDesc: 'SSE live stream with recent query records',
    logSearchPlaceholder: 'Search domain / client / answer',
    protocolFilterAria: 'Protocol filter',
    rcodeFilterAria: 'RCODE filter',
    allProtocols: 'All protocols',
    allRcodes: 'All rcodes',
    searchButton: 'Search',
    resetButton: 'Reset',
    prevButton: 'Prev',
    nextButton: 'Next',
    pageInfo: 'Page {page} / {totalPages} · {total} results',
    noQueryLogs: 'No query logs yet',
    logHeadTime: 'Time',
    logHeadProto: 'Proto',
    logHeadSource: 'Source',
    logHeadUpstream: 'Upstream',
    logHeadQName: 'QName',
    logHeadAnswer: 'Answer',
    logHeadTTL: 'TTL',
    logHeadType: 'Type',
    logHeadRcode: 'RCODE',
    logHeadLatency: 'Latency',
    logHeadClient: 'Client',
    logHeadDetail: 'Detail',
    sourceCache: 'Cache',
    sourceUpstream: 'Upstream',
    detailsButton: 'Details',
    queryDetailTitle: 'Query Detail',
    closeDetailAria: 'Close detail',
    detailProtocol: 'Protocol',
    detailQType: 'QType',
    detailQClass: 'QClass',
    detailAnswers: 'Answers',
    detailError: 'Error',
    upstreamServersTitle: 'Upstream Servers',
    upstreamServersDesc: 'Status, latency, success rate, and in-flight requests',
    upstreamGroupCount: '{count} upstreams',
    noUpstreams: 'No tagged upstream statistics yet',
    testQueryButton: 'Test',
    testQueryTitle: 'Test Query',
    testQueryDomain: 'Domain',
    testQueryType: 'Type',
    testQueryResult: 'Result',
    testQueryUpstream: 'Upstream',
    testQueryTime: 'Time',
    testQueryAnswers: 'Answers',
    testQueryError: 'Error',
    testQueryWorking: 'Querying...',
    testQuerySuccess: 'Query successful',
    testQueryFailed: 'Query failed',
    statQueries: 'Queries',
    statSuccessRate: 'Success Rate',
    statLastLatency: 'Last Latency',
    statInFlight: 'In Flight',
    cachePluginsTitle: 'Cache Plugins',
    cachePluginsDesc: 'Cache hits and capacity',
    noCacheStats: 'No cache plugin statistics yet',
    statHits: 'Hits',
    statLazyHits: 'Lazy Hits',
    statSize: 'Size',
    configurationTitle: 'Configuration',
    configurationDesc: 'Edit the active main config file. A .bak backup is created before saving.',
    loadButton: 'Load',
    validateButton: 'Validate',
    saveButton: 'Save',
    reloadButton: 'Reload',
    configInitialMessage: 'Reload is a placeholder in this first version and does not perform a real reload yet.',
    configLoaded: 'Config loaded.',
    configValid: 'Config is valid.',
    configInvalid: 'Config validation failed.',
    configSaved: 'Saved. Backup: {backup}',
    configSavedToast: 'Config saved',
    reloadNotImplemented: 'Reload is not implemented yet.',
  },
  zh: {
    sidebarNavAria: '主导航',
    pageNavAria: '页面',
    languageToggleAria: '切换语言',
    languageToggle: 'English',
    navDashboard: '仪表盘',
    navLogs: '查询日志',
    navUpstreams: '上游服务',
    navConfig: '配置',
    storageLabel: '存储',
    storageMemoryOnly: '仅内存',
    storageRedisError: 'Redis 错误：{queue} 条排队',
    storageRedisActive: 'Redis 已启用：{queue} 条排队',
    eyebrow: 'DNS 可观测性',
    liveStatus: '实时',
    refreshButton: '刷新',
    logoutButton: '退出',
    loginHeroKicker: '安全 DNS 控制台',
    loginHeroText: '登录后查看实时查询、上游状态和配置。',
    loginEyebrow: 'WebUI 访问',
    loginTitle: '登录',
    loginCopy: '使用 mosdns 中配置的 WebUI 用户名和密码。',
    loginUsername: '用户名',
    loginPassword: '密码',
    loginButton: '登录',
    loginWorking: '正在登录...',
    loginFailed: '用户名或密码错误。',
    loginRequired: '请先登录。',
    viewDashboard: '仪表盘',
    viewLogs: '查询日志',
    viewUpstreams: '上游服务',
    viewConfig: '配置编辑器',
    metricTotalQueries: '总查询数',
    metricTotalQueriesHint: '累计处理请求',
    metricQps: 'QPS',
    metricQpsHint: '最近 10 秒平均',
    metricAvgLatency: '平均延迟',
    metricAvgLatencyHint: '端到端查询耗时',
    metricCacheHitRate: '缓存命中率',
    cacheSizeHint: '{count} 条记录',
    refreshCacheTitle: '刷新域名缓存',
    refreshCacheDesc: '强制移除指定域名缓存，并立即重新查询写入缓存',
    cacheRefreshTypeAria: '缓存刷新类型',
    cacheRefreshEntryAria: '刷新入口插件',
    refreshCacheButton: '刷新缓存',
    refreshCacheInitial: '输入域名后点击刷新。',
    refreshCacheWorking: '正在刷新缓存...',
    refreshCacheSuccess: '缓存已刷新',
    refreshCacheRequired: '请先输入域名',
    refreshRemovedSummary: '已移除 {count} 个缓存键 · 入口 {entry}',
    noRefreshResult: '暂无刷新结果',
    hourlyQueriesTitle: '每小时查询',
    hourlyQueriesDesc: '最近小时级请求分布',
    hourlyChartAria: '每小时查询次数',
    noHourlyData: '暂无小时查询数据',
    domainTopTitle: '域名 Top 10',
    domainTopDesc: '请求最频繁的域名',
    noDomainData: '暂无域名数据',
    topClientsTitle: '客户端排行',
    topClientsDesc: '客户端请求量排行',
    noClientData: '暂无客户端数据',
    realtimeTailTitle: '实时查询尾部',
    realtimeTailDesc: '最近 10 条请求',
    openLogs: '打开日志',
    queryLogsTitle: '实时查询日志',
    queryLogsDesc: 'SSE 实时推送，保留最近查询记录',
    logSearchPlaceholder: '搜索域名 / 客户端 / 响应地址',
    protocolFilterAria: '协议筛选',
    rcodeFilterAria: '响应码筛选',
    allProtocols: '所有协议',
    allRcodes: '所有响应码',
    searchButton: '搜索',
    resetButton: '重置',
    prevButton: '上一页',
    nextButton: '下一页',
    pageInfo: '第 {page} / {totalPages} 页 · {total} 条结果',
    noQueryLogs: '暂无查询日志',
    logHeadTime: '时间',
    logHeadProto: '协议',
    logHeadSource: '来源',
    logHeadUpstream: '上游',
    logHeadQName: '域名',
    logHeadAnswer: '响应',
    logHeadTTL: 'TTL',
    logHeadType: '类型',
    logHeadRcode: '响应码',
    logHeadLatency: '延迟',
    logHeadClient: '客户端',
    logHeadDetail: '详情',
    sourceCache: '缓存',
    sourceUpstream: '上游',
    detailsButton: '详情',
    queryDetailTitle: '查询详情',
    closeDetailAria: '关闭详情',
    detailProtocol: '协议',
    detailQType: '查询类型',
    detailQClass: '查询类',
    detailAnswers: '响应',
    detailError: '错误',
    upstreamServersTitle: '上游服务器',
    upstreamServersDesc: '状态、延迟、成功率和并发请求',
    upstreamGroupCount: '{count} 个上游',
    noUpstreams: '暂无带标签的上游统计',
    testQueryButton: '测试',
    testQueryTitle: '测试查询',
    testQueryDomain: '域名',
    testQueryType: '类型',
    testQueryResult: '结果',
    testQueryUpstream: '上游',
    testQueryTime: '时间',
    testQueryAnswers: '响应',
    testQueryError: '错误',
    testQueryWorking: '正在查询...',
    testQuerySuccess: '查询成功',
    testQueryFailed: '查询失败',
    statQueries: '查询数',
    statSuccessRate: '成功率',
    statLastLatency: '最近延迟',
    statInFlight: '进行中',
    cachePluginsTitle: '缓存插件',
    cachePluginsDesc: '缓存命中和容量',
    noCacheStats: '暂无缓存插件统计',
    statHits: '命中数',
    statLazyHits: '懒缓存命中',
    statSize: '容量',
    configurationTitle: '配置',
    configurationDesc: '编辑当前主配置文件。保存前会自动创建 .bak 备份。',
    loadButton: '加载',
    validateButton: '校验',
    saveButton: '保存',
    reloadButton: '重载',
    configInitialMessage: '配置重载按钮为占位功能，初版暂不执行真实 reload。',
    configLoaded: '配置已加载。',
    configValid: '配置有效。',
    configInvalid: '配置校验失败。',
    configSaved: '已保存。备份：{backup}',
    configSavedToast: '配置已保存',
    reloadNotImplemented: '暂未实现重载功能。',
  },
};

const state = {
  logs: [],
  recentLogs: [],
  maxLogs: 300,
  currentView: 'dashboard',
  protocolFilter: '',
  rcodeFilter: '',
  searchText: '',
  page: 1,
  pageSize: 50,
  totalLogs: 0,
  language: ['en', 'zh'].includes(localStorage.getItem('webui-language')) ? localStorage.getItem('webui-language') : 'en',
  cacheSize: 0,
  storageStatus: null,
  authenticated: false,
  logSource: null,
  refreshTimer: null,
};

const $ = (id) => document.getElementById(id);

function t(key, params = {}) {
  const template = translations[state.language]?.[key] ?? translations.en[key] ?? key;
  return Object.entries(params).reduce((text, [name, value]) => text.replaceAll(`{${name}}`, value), template);
}

function applyLanguage() {
  document.documentElement.lang = state.language === 'zh' ? 'zh-CN' : 'en';
  document.querySelectorAll('[data-i18n]').forEach((el) => {
    el.textContent = t(el.dataset.i18n);
  });
  document.querySelectorAll('[data-i18n-attr]').forEach((el) => {
    el.dataset.i18nAttr.split(',').forEach((pair) => {
      const [attr, key] = pair.split(':');
      el.setAttribute(attr, t(key));
    });
  });
  $('languageToggle').textContent = t('languageToggle');
  $('loginLanguageToggle').textContent = t('languageToggle');
  $('cacheSizeHint').textContent = t('cacheSizeHint', { count: fmtNumber(state.cacheSize) });
  if (state.storageStatus) renderStorageStatus(state.storageStatus);
  setView(state.currentView, false);
  rerenderLogs();
  renderLogPager();
}

function setLanguage(language) {
  state.language = translations[language] ? language : 'en';
  localStorage.setItem('webui-language', state.language);
  applyLanguage();
}

function fmtNumber(value) {
  const locale = state.language === 'zh' ? 'zh-CN' : 'en-US';
  return new Intl.NumberFormat(locale).format(Number(value || 0));
}

function fmtTime(value) {
  if (!value) return '-';
  return new Date(value).toLocaleTimeString([], { hour12: false });
}

function fmtDateTime(value) {
  if (!value) return '-';
  return new Date(value).toLocaleString([], { hour12: false });
}

function qtypeName(type) {
  const names = { 1: 'A', 2: 'NS', 5: 'CNAME', 12: 'PTR', 15: 'MX', 16: 'TXT', 28: 'AAAA', 33: 'SRV', 65: 'HTTPS' };
  return names[type] || String(type || '-');
}

function rcodeName(rcode) {
  if (rcode === null || rcode === undefined || rcode === '') return '-';
  const names = {
    0: 'NOERROR',
    1: 'FORMERR',
    2: 'SERVFAIL',
    3: 'NXDOMAIN',
    4: 'NOTIMP',
    5: 'REFUSED',
    6: 'YXDOMAIN',
    7: 'YXRRSET',
    8: 'NXRRSET',
    9: 'NOTAUTH',
    10: 'NOTZONE',
    11: 'DSOTYPENI',
    16: 'BADVERS',
    17: 'BADKEY',
    18: 'BADTIME',
    19: 'BADMODE',
    20: 'BADNAME',
    21: 'BADALG',
    22: 'BADTRUNC',
    23: 'BADCOOKIE',
  };
  return names[Number(rcode)] || String(rcode);
}

function rcodeWithCode(rcode) {
  const name = rcodeName(rcode);
  if (name === '-') return name;
  return `${name} (${rcode})`;
}

function protocolName(protocol) {
  const value = String(protocol || '').toLowerCase();
  const names = { udp: 'UDP', tcp: 'TCP', dot: 'DoT', doh: 'DoH', doq: 'DoQ' };
  return names[value] || value || '-';
}

function renderAnswer(log) {
  if (!log.answers || log.answers.length === 0) return '-';
  return log.answers.join(', ');
}

function renderTTL(log) {
  if (!log.ttls || log.ttls.length === 0) return '-';
  return log.ttls.map((ttl) => `${ttl}s`).join(', ');
}

async function request(path, options = {}) {
  const response = await fetch(`${API}${path}`, options);
  if (!response.ok) {
    const text = await response.text();
    const error = new Error(text || `${response.status} ${response.statusText}`);
    error.status = response.status;
    if (response.status === 401 && path !== '/session') showLogin(t('loginRequired'));
    throw error;
  }
  const type = response.headers.get('content-type') || '';
  if (type.includes('application/json')) return response.json();
  return response.text();
}

function showLogin(message = '') {
  state.authenticated = false;
  stopRealtimeUpdates();
  $('appShell').hidden = true;
  $('loginPage').hidden = false;
  $('loginMessage').textContent = message;
  setTimeout(() => $('loginUsername').focus(), 0);
}

function showApp() {
  state.authenticated = true;
  $('loginPage').hidden = true;
  $('appShell').hidden = false;
  $('loginMessage').textContent = '';
}

async function checkSession() {
  const session = await request('/session');
  return Boolean(session.authenticated);
}

async function login() {
  const button = $('loginButton');
  button.disabled = true;
  $('loginMessage').textContent = t('loginWorking');
  try {
    await request('/session', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        username: $('loginUsername').value,
        password: $('loginPassword').value,
      }),
    });
    $('loginPassword').value = '';
    showApp();
    await startApp();
  } catch (error) {
    $('loginMessage').textContent = error.status === 401 ? t('loginFailed') : error.message.trim();
  } finally {
    button.disabled = false;
  }
}

async function logout() {
  try {
    await request('/session', { method: 'DELETE' });
  } finally {
    showLogin('');
  }
}

async function startApp() {
  setView((location.hash || '#dashboard').slice(1));
  await refreshAll();
  connectLogStream();
  clearInterval(state.refreshTimer);
  state.refreshTimer = setInterval(refreshAll, 5000);
}

function stopRealtimeUpdates() {
  if (state.logSource) {
    state.logSource.close();
    state.logSource = null;
  }
  clearInterval(state.refreshTimer);
  state.refreshTimer = null;
}

function showToast(message) {
  const toast = $('toast');
  toast.textContent = message;
  toast.classList.add('is-visible');
  clearTimeout(showToast.timer);
  showToast.timer = setTimeout(() => toast.classList.remove('is-visible'), 2600);
}

function setDynamicText(id, text) {
  const el = $(id);
  el.removeAttribute('data-i18n');
  el.textContent = text;
}

function empty(label) {
  return `<div class="empty-state">${label}</div>`;
}

function setView(view, updateHash = true) {
  state.currentView = view;
  document.querySelectorAll('.view').forEach((el) => el.classList.remove('is-visible'));
  document.querySelectorAll('.nav-item, .link-button[data-view]').forEach((el) => el.classList.toggle('is-active', el.dataset.view === view));
  const viewMap = {
    dashboard: ['dashboardView', t('viewDashboard')],
    logs: ['logsView', t('viewLogs')],
    upstreams: ['upstreamsView', t('viewUpstreams')],
    config: ['configView', t('viewConfig')],
  };
  const [id, title] = viewMap[view] || viewMap.dashboard;
  $(id).classList.add('is-visible');
  $('viewTitle').textContent = title;
  if (updateHash) location.hash = view;
}

async function loadOverview() {
  const [overview, cacheStats] = await Promise.all([request('/overview'), request('/cache')]);
  $('totalQueries').textContent = fmtNumber(overview.total_queries);
  $('qps').textContent = Number(overview.qps || 0).toFixed(2);
  $('avgLatency').textContent = Number(overview.avg_latency_ms || 0).toFixed(1);

  const cacheQuery = cacheStats.reduce((sum, item) => sum + (item.query_total || 0), 0);
  const cacheHit = cacheStats.reduce((sum, item) => sum + (item.hit_total || 0), 0);
  const cacheSize = cacheStats.reduce((sum, item) => sum + (item.size || 0), 0);
  state.cacheSize = cacheSize;
  const hitRate = cacheQuery > 0 ? (cacheHit / cacheQuery) * 100 : 0;
  $('cacheHitRate').textContent = hitRate.toFixed(1);
  $('cacheSizeHint').textContent = t('cacheSizeHint', { count: fmtNumber(cacheSize) });
}

async function loadDashboard() {
  const [hourly, domains, clients, logs] = await Promise.all([
    request('/stats/hourly'),
    request('/stats/domains/top?limit=10'),
    request('/stats/clients?limit=10'),
    request('/query-log/recent?limit=10'),
  ]);
  state.recentLogs = logs || [];
  renderBars(hourly);
  renderRanks('domainTop', domains, t('noDomainData'));
  renderRanks('clientTop', clients, t('noClientData'));
  renderLogs('recentLogTable', state.recentLogs, true);
}

function renderBars(points) {
  const container = $('hourlyChart');
  if (!points || points.length === 0) {
    container.innerHTML = empty(t('noHourlyData'));
    return;
  }
  const max = Math.max(...points.map((p) => p.count), 1);
  container.innerHTML = points.map((p) => {
    const height = Math.max(6, (p.count / max) * 100);
    const label = new Date(p.time).getHours().toString().padStart(2, '0');
    return `<div class="bar" title="${label}:00 ${p.count}" data-label="${label}" style="height:${height}%"></div>`;
  }).join('');
}

function renderRanks(id, rows, fallback) {
  const container = $(id);
  if (!rows || rows.length === 0) {
    container.innerHTML = empty(fallback);
    return;
  }
  container.innerHTML = rows.map((row) => `
    <div class="rank-item">
      <span class="rank-name" title="${escapeHtml(row.name)}">${escapeHtml(row.name)}</span>
      <span class="rank-count">${fmtNumber(row.count)}</span>
    </div>
  `).join('');
}

function renderRefreshResult(result) {
  const rows = (result.results || []).map((item) => {
    const answers = item.answers && item.answers.length > 0 ? item.answers.join(', ') : '-';
    const ttls = item.ttls && item.ttls.length > 0 ? item.ttls.map((ttl) => `${ttl}s`).join(', ') : '-';
    const status = item.error ? item.error : rcodeWithCode(item.rcode);
    return `
      <div class="refresh-result-row">
        <span>${escapeHtml(item.qtype || '-')}</span>
        <span>${escapeHtml(status)}</span>
        <span title="${escapeHtml(answers)}">${escapeHtml(answers)}</span>
        <span title="${escapeHtml(ttls)}">${escapeHtml(ttls)}</span>
      </div>
    `;
  }).join('');
  return `
    <div class="refresh-result-summary">
      <strong>${escapeHtml(result.domain || '-')}</strong>
      <span>${escapeHtml(t('refreshRemovedSummary', { count: fmtNumber(result.removed), entry: result.entry || '-' }))}</span>
    </div>
    ${rows || `<div class="empty-state">${t('noRefreshResult')}</div>`}
  `;
}

async function refreshDomainCache() {
  const domain = $('cacheRefreshDomain').value.trim();
  if (!domain) {
    showToast(t('refreshCacheRequired'));
    $('cacheRefreshDomain').focus();
    return;
  }
  const button = $('cacheRefreshButton');
  button.disabled = true;
  setDynamicText('cacheRefreshResult', t('refreshCacheWorking'));
  try {
    const result = await request('/cache/refresh', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        domain,
        qtype: $('cacheRefreshType').value,
        entry: $('cacheRefreshEntry').value.trim(),
      }),
    });
    $('cacheRefreshResult').removeAttribute('data-i18n');
    $('cacheRefreshResult').innerHTML = renderRefreshResult(result);
    showToast(t('refreshCacheSuccess'));
    await Promise.all([loadOverview(), loadDashboard()]);
  } finally {
    button.disabled = false;
  }
}

function renderLogs(id, logs, compact = false) {
  const container = $(id);
  if (!logs || logs.length === 0) {
    container.innerHTML = empty(t('noQueryLogs'));
    return;
  }
  const rows = logs.map((log, index) => logRow(log, compact, index)).join('');
  container.innerHTML = compact ? rows : `${logHeader()}${rows}`;
}

function logHeader() {
  return `
    <div class="log-row log-head" aria-hidden="true">
      <span>${t('logHeadTime')}</span>
      <span>${t('logHeadProto')}</span>
      <span>${t('logHeadSource')}</span>
      <span>${t('logHeadUpstream')}</span>
      <span>${t('logHeadQName')}</span>
      <span>${t('logHeadAnswer')}</span>
      <span>${t('logHeadTTL')}</span>
      <span>${t('logHeadType')}</span>
      <span>${t('logHeadRcode')}</span>
      <span>${t('logHeadLatency')}</span>
      <span>${t('logHeadClient')}</span>
      <span>${t('logHeadDetail')}</span>
    </div>
  `;
}

function logRow(log, compact, index) {
  const error = log.rcode !== 0 || log.error;
  const source = log.cache_hit ? t('sourceCache') : t('sourceUpstream');
  const upstream = log.upstream_tag || log.upstream_addr || '-';
  return `
    <div class="log-row">
      <span>${compact ? fmtTime(log.time) : fmtDateTime(log.time)}</span>
      <span class="protocol-badge">${protocolName(log.protocol)}</span>
      ${compact ? '' : `<span class="source-badge ${log.cache_hit ? 'is-cache' : ''}">${source}</span>`}
      ${compact ? '' : `<span class="upstream-info" title="${escapeHtml(log.upstream_addr || '')}">${escapeHtml(upstream)}</span>`}
      <span class="log-qname" title="${escapeHtml(log.qname)}">${escapeHtml(log.qname || '-')}</span>
      <span class="log-answer" title="${escapeHtml(renderAnswer(log))}">${escapeHtml(renderAnswer(log))}</span>
      <span class="log-ttl" title="${escapeHtml(renderTTL(log))}">${escapeHtml(renderTTL(log))}</span>
      <span>${qtypeName(log.qtype)}</span>
      <span class="rcode ${error ? 'is-error' : ''}" title="${escapeHtml(rcodeWithCode(log.rcode))}">${escapeHtml(rcodeName(log.rcode))}</span>
      <span>${Number(log.elapsed_ms || 0)}ms</span>
      ${compact ? '' : `<span title="${escapeHtml(log.client || '')}">${escapeHtml(log.client || '-')}</span>`}
      ${compact ? '' : `<button class="detail-button" type="button" data-log-index="${index}">${t('detailsButton')}</button>`}
    </div>
  `;
}

function openLogDetail(index) {
  const log = state.logs[Number(index)];
  if (!log) return;
  const source = log.cache_hit ? t('sourceCache') : t('sourceUpstream');
  const upstream = log.upstream_tag || log.upstream_addr || '-';
  const pairs = [
    [t('logHeadTime'), fmtDateTime(log.time)],
    [t('detailProtocol'), protocolName(log.protocol)],
    [t('logHeadSource'), source],
    [t('logHeadUpstream'), upstream],
    [t('logHeadClient'), log.client || '-'],
    ['QName', log.qname || '-'],
    [t('detailQType'), qtypeName(log.qtype)],
    [t('detailQClass'), log.qclass || '-'],
    ['RCODE', rcodeWithCode(log.rcode)],
    [t('logHeadLatency'), `${Number(log.elapsed_ms || 0)}ms`],
    ['TTL', renderTTL(log)],
    [t('detailAnswers'), renderAnswer(log)],
    [t('detailError'), log.error || '-'],
  ];
  $('logDetailBody').innerHTML = pairs.map(([label, value]) => `
    <div class="detail-item">
      <span>${escapeHtml(label)}</span>
      <strong title="${escapeHtml(value)}">${escapeHtml(value)}</strong>
    </div>
  `).join('');
  $('logDetailModal').hidden = false;
}

function closeLogDetail() {
  $('logDetailModal').hidden = true;
}

async function loadLogs() {
  const params = new URLSearchParams({
    limit: String(state.pageSize),
    offset: String((state.page - 1) * state.pageSize),
  });
  if (state.searchText) params.set('search', state.searchText);
  if (state.protocolFilter) params.set('protocol', state.protocolFilter);
  if (state.rcodeFilter) params.set('rcode', state.rcodeFilter);
  const page = await request(`/query-log?${params}`);
  state.logs = page.rows || [];
  state.totalLogs = page.total || 0;
  renderLogs('logTable', state.logs);
  renderLogPager();
}

function renderLogPager() {
  const totalPages = Math.max(1, Math.ceil(state.totalLogs / state.pageSize));
  if (state.page > totalPages) state.page = totalPages;
  $('logPageInfo').textContent = t('pageInfo', { page: state.page, totalPages, total: fmtNumber(state.totalLogs) });
  $('prevLogPage').disabled = state.page <= 1;
  $('nextLogPage').disabled = state.page >= totalPages;
}

function hasLogFilters() {
  return Boolean(state.searchText || state.protocolFilter || state.rcodeFilter);
}

function applyLogSearch() {
  state.searchText = $('logSearch').value.trim();
  state.protocolFilter = $('protocolFilter').value;
  state.rcodeFilter = $('rcodeFilter').value;
  state.page = 1;
  loadLogs().catch((error) => showToast(error.message));
}

function resetLogFilters() {
  $('logSearch').value = '';
  $('protocolFilter').value = '';
  $('rcodeFilter').value = '';
  state.searchText = '';
  state.protocolFilter = '';
  state.rcodeFilter = '';
  state.page = 1;
  loadLogs().catch((error) => showToast(error.message));
}

async function loadUpstreams() {
  const [upstreams, caches] = await Promise.all([request('/upstreams'), request('/cache')]);
  const grid = $('upstreamGrid');
  if (!upstreams || upstreams.length === 0) {
    grid.innerHTML = empty(t('noUpstreams'));
  } else {
    grid.innerHTML = renderUpstreamGroups(upstreams);
  }

  const list = $('cacheList');
  if (!caches || caches.length === 0) {
    list.innerHTML = empty(t('noCacheStats'));
  } else {
    list.innerHTML = caches.map((c) => `
      <div class="cache-item">
        <div class="upstream-title"><span>${escapeHtml(c.tag || 'cache')}</span><span>${((c.hit_rate || 0) * 100).toFixed(1)}%</span></div>
        <div class="stat-pairs">
          <div class="stat-pair"><span>${t('statQueries')}</span><strong>${fmtNumber(c.query_total)}</strong></div>
          <div class="stat-pair"><span>${t('statHits')}</span><strong>${fmtNumber(c.hit_total)}</strong></div>
          <div class="stat-pair"><span>${t('statLazyHits')}</span><strong>${fmtNumber(c.lazy_hit_total)}</strong></div>
          <div class="stat-pair"><span>${t('statSize')}</span><strong>${fmtNumber(c.size)}</strong></div>
        </div>
      </div>
    `).join('');
  }
}

function renderUpstreamGroups(upstreams) {
  const sorted = [...upstreams].sort((a, b) => String(a.forward_tag || '').localeCompare(String(b.forward_tag || '')));
  const groups = [];
  for (const upstream of sorted) {
    const groupTag = upstream.forward_tag || 'forward';
    let group = groups[groups.length - 1];
    if (!group || group.tag !== groupTag) {
      group = { tag: groupTag, items: [] };
      groups.push(group);
    }
    group.items.push(upstream);
  }
  return groups.map((group) => `
    <section class="upstream-group">
      <div class="upstream-group-header">
        <span>${escapeHtml(group.tag)}</span>
        <small>${escapeHtml(t('upstreamGroupCount', { count: fmtNumber(group.items.length) }))}</small>
      </div>
      <div class="upstream-group-grid">
        ${group.items.map(renderUpstreamCard).join('')}
      </div>
    </section>
  `).join('');
}

function renderUpstreamCard(u) {
  return `
    <div class="upstream-card">
      <div class="upstream-title">
        <span>${escapeHtml(u.tag || 'upstream')}</span>
        <span class="status ${escapeHtml(u.status || 'unknown')}">${escapeHtml(u.status || 'unknown')}</span>
      </div>
      <div class="upstream-addr" title="${escapeHtml(u.addr)}">${escapeHtml(u.addr || '-')}</div>
      <div class="stat-pairs">
        <div class="stat-pair"><span>${t('statQueries')}</span><strong>${fmtNumber(u.query_total)}</strong></div>
        <div class="stat-pair"><span>${t('statSuccessRate')}</span><strong>${((u.success_rate || 0) * 100).toFixed(1)}%</strong></div>
        <div class="stat-pair"><span>${t('statLastLatency')}</span><strong>${u.last_latency_ms || 0}ms</strong></div>
        <div class="stat-pair"><span>${t('statInFlight')}</span><strong>${u.inflight || 0}</strong></div>
      </div>
      <button class="button ghost test-upstream-button" type="button" data-upstream-tag="${escapeHtml(u.tag || '')}" data-upstream-addr="${escapeHtml(u.addr || '')}">${t('testQueryButton')}</button>
    </div>
  `;
}

async function loadConfig() {
  const text = await request('/config');
  $('configEditor').value = text;
  setDynamicText('configMessage', t('configLoaded'));
}

async function testUpstreamQuery(tag, addr) {
  const domain = prompt(t('testQueryDomain') + ':', 'example.com');
  if (!domain) return;
  
  showToast(t('testQueryWorking'));
  try {
    const result = await request('/query-test', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        domain: domain,
        qtype: 'A',
        upstream_tag: tag,
        upstream_addr: addr,
      }),
    });
    
    const answers = result.answers && result.answers.length > 0 ? result.answers.join(', ') : '-';
    const message = `${t('testQueryUpstream')}: ${tag || addr}\n${t('testQueryTime')}: ${result.elapsed_ms || 0}ms\n${t('testQueryAnswers')}: ${answers}`;
    alert(message);
    showToast(t('testQuerySuccess'));
  } catch (error) {
    showToast(t('testQueryFailed') + ': ' + error.message);
  }
}

async function validateConfig() {
  const result = await request('/config/validate', { method: 'POST', body: $('configEditor').value });
  setDynamicText('configMessage', result.ok ? t('configValid') : t('configInvalid'));
  showToast(result.ok ? t('configValid') : t('configInvalid'));
}

async function saveConfig() {
  const body = $('configEditor').value;
  const result = await request('/config', { method: 'PUT', body });
  setDynamicText('configMessage', t('configSaved', { backup: result.backup || '-' }));
  showToast(t('configSavedToast'));
}

async function reloadConfig() {
  try {
    await request('/config/reload', { method: 'POST' });
  } catch (error) {
    setDynamicText('configMessage', t('reloadNotImplemented'));
    showToast(t('reloadNotImplemented'));
  }
}

async function loadStorageStatus() {
  const status = await request('/storage/status');
  state.storageStatus = status;
  renderStorageStatus(status);
}

function renderStorageStatus(status) {
  if (!status.enabled) {
    $('storageStatus').textContent = t('storageMemoryOnly');
    return;
  }
  $('storageStatus').textContent = status.last_error
    ? t('storageRedisError', { queue: fmtNumber(status.queue_len) })
    : t('storageRedisActive', { queue: fmtNumber(status.queue_len) });
}

async function refreshAll() {
  try {
    await loadOverview();
    await Promise.all([loadDashboard(), loadStorageStatus()]);
    if (state.currentView === 'logs') await loadLogs();
    if (state.currentView === 'upstreams') await loadUpstreams();
  } catch (error) {
    showToast(error.message.trim());
  }
}

function connectLogStream() {
  if (state.logSource) state.logSource.close();
  const live = $('liveStatus');
  const source = new EventSource(`${API}/query-log/stream`);
  state.logSource = source;
  source.onopen = () => live.classList.remove('is-offline');
  source.onerror = () => live.classList.add('is-offline');
  source.onmessage = (event) => {
    const log = JSON.parse(event.data);
    state.recentLogs.unshift(log);
    state.recentLogs = state.recentLogs.slice(0, 10);
    renderLogs('recentLogTable', state.recentLogs, true);
    if (state.currentView === 'logs' && state.page === 1 && !hasLogFilters()) {
      state.logs.unshift(log);
      state.logs = state.logs.slice(0, state.pageSize);
      state.totalLogs++;
      renderLogs('logTable', state.logs);
      renderLogPager();
    }
    loadOverview().catch(() => {});
  };
}

function rerenderLogs() {
  renderLogs('logTable', state.logs);
  renderLogs('recentLogTable', state.recentLogs, true);
  renderLogPager();
}

function escapeHtml(value) {
  return String(value ?? '')
    .replaceAll('&', '&amp;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;')
    .replaceAll('"', '&quot;')
    .replaceAll("'", '&#039;');
}

document.addEventListener('click', (event) => {
  const detailButton = event.target.closest('[data-log-index]');
  if (detailButton) {
    openLogDetail(detailButton.dataset.logIndex);
    return;
  }
  const testButton = event.target.closest('.test-upstream-button');
  if (testButton) {
    testUpstreamQuery(testButton.dataset.upstreamTag, testButton.dataset.upstreamAddr);
    return;
  }
  const link = event.target.closest('[data-view]');
  if (link) {
    event.preventDefault();
    setView(link.dataset.view);
    if (link.dataset.view === 'logs') loadLogs().catch((error) => showToast(error.message));
    if (link.dataset.view === 'upstreams') loadUpstreams().catch((error) => showToast(error.message));
    if (link.dataset.view === 'config' && !$('configEditor').value) loadConfig().catch((error) => showToast(error.message));
  }
});

$('logDetailModal').addEventListener('click', (event) => {
  if (event.target === $('logDetailModal')) closeLogDetail();
});
$('closeLogDetail').addEventListener('click', closeLogDetail);
document.addEventListener('keydown', (event) => {
  if (event.key === 'Escape' && !$('logDetailModal').hidden) closeLogDetail();
});

$('refreshButton').addEventListener('click', () => refreshAll());
$('languageToggle').addEventListener('click', () => {
  setLanguage(state.language === 'en' ? 'zh' : 'en');
  refreshAll();
});
$('loginLanguageToggle').addEventListener('click', () => {
  setLanguage(state.language === 'en' ? 'zh' : 'en');
});
$('loginForm').addEventListener('submit', (event) => {
  event.preventDefault();
  login();
});
$('logoutButton').addEventListener('click', () => logout());
$('cacheRefreshButton').addEventListener('click', () => refreshDomainCache().catch((error) => {
  setDynamicText('cacheRefreshResult', error.message);
  showToast(error.message);
}));
$('cacheRefreshDomain').addEventListener('keydown', (event) => { if (event.key === 'Enter') refreshDomainCache().catch((error) => showToast(error.message)); });
$('resetLogButton').addEventListener('click', resetLogFilters);
$('searchLogButton').addEventListener('click', applyLogSearch);
$('logSearch').addEventListener('keydown', (event) => { if (event.key === 'Enter') applyLogSearch(); });
$('prevLogPage').addEventListener('click', () => {
  if (state.page <= 1) return;
  state.page--;
  loadLogs().catch((error) => showToast(error.message));
});
$('nextLogPage').addEventListener('click', () => {
  if (state.page * state.pageSize >= state.totalLogs) return;
  state.page++;
  loadLogs().catch((error) => showToast(error.message));
});
$('loadConfigButton').addEventListener('click', () => loadConfig().catch((error) => showToast(error.message)));
$('validateConfigButton').addEventListener('click', () => validateConfig().catch((error) => showToast(error.message)));
$('saveConfigButton').addEventListener('click', () => saveConfig().catch((error) => showToast(error.message)));
$('reloadConfigButton').addEventListener('click', () => reloadConfig());

applyLanguage();
checkSession()
  .then((authenticated) => {
    if (!authenticated) {
      showLogin('');
      return;
    }
    showApp();
    startApp();
  })
  .catch(() => showLogin(''));
